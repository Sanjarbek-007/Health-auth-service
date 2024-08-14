package service

import (
	pb "auth-service/genproto/user"
	"auth-service/logs"
	"auth-service/storage"
	"context"
	"fmt"
	"log/slog"
)

type UserService struct {
	pb.UnimplementedUsersServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewUserService(s storage.IStorage) *UserService {
	return &UserService{
		storage: s,
		logger:  logs.NewLogger(),
	}
}

func (s *UserService) GetUserById(ctx context.Context, req *pb.UserId) (*pb.User, error) {
	resp, err := s.storage.User().GetByUserID(ctx, req.UserId)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error retrieving id information: %v", err))
		return nil, err
	}
	return resp, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.Void, error) {
	_, err := s.storage.User().UpdateProfile(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Error Update user: %v", err))
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.UserId) (*pb.Void, error) {
	err := s.storage.User().DeleteUser(ctx, req.UserId)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error delete user: %v", err))
		return nil, err
	}
	return &pb.Void{}, nil
}
