package service

import (
	pb "auth-service/genproto/user"
	"auth-service/logs"
	"auth-service/storage"
	"context"
	"fmt"
	"log/slog"
)

type NoteService struct {
	pb.UnimplementedUsersServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewNoteService(s storage.IStorage) *UserService {
	return &UserService{
		storage: s,
		logger:  logs.NewLogger(),
	}
}

func (s *NoteService) GetAllNotifications(ctx context.Context, req *pb.GetNotificationsReq) (*pb.GetNotificationsResponse, error) {
	s.logger.Info("GetAllNotifications rpc method is working")
	resp, err := s.storage.Notifications().GetAllNotifications(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting notifications: %v", err))
		return nil, err
	}
	s.logger.Info("GetAllNotifications rpc method finished")
	return resp, nil
}

func (s *NoteService) GetAndMarkNotificationAsRead(ctx context.Context, req *pb.GetAndMarkNotificationAsReadReq) (*pb.GetAndMarkNotificationAsReadRes, error) {
	s.logger.Info("GetAndMarkNotificationAsRead rpc method is working")
	resp, err := s.storage.Notifications().GetAndMarkNotificationAsRead(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error getting and marking notification as read: %v", err))
		return nil, err
	}
	s.logger.Info("GetAndMarkNotificationAsRead rpc method finished")
	return resp, nil
}
