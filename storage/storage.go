package storage

import (
	pb "auth-service/genproto/user"
	"auth-service/models"
	"context"
)

type IStorage interface {
	User() IUserStorage
	Notifications() INotificationStorage
	Close()
}

type IUserStorage interface {
	CreateUser(context.Context, *pb.RegisterReq) (*pb.RegisterRes, error)
	GetUserDetails(context.Context, string) (*models.UserDetails, error)
	GetUserProfile(context.Context, *pb.GetProfileReq) (*pb.GetProfileRes, error)
	UpdateProfile(context.Context, *pb.UpdateProfileReq) (*pb.Void, error)
	UpdatePassword(context.Context, *pb.UpdatePasswordReq) error
	GetUSerByEmail(context.Context, *pb.GetUSerByEmailReq) (*pb.FilterUsers, error)
	ChangePassword(ctx context.Context, userID, hashedPassword string) error
	GetByUserID(ctx context.Context, userID string) (*pb.User, error)
	DeleteUser(ctx context.Context, userID string) error
}

type INotificationStorage interface {
	CreateNotifications(context.Context, *pb.CreateNotificationsReq) (*pb.CreateNotificationsRes, error)
	GetAllNotifications(context.Context, *pb.GetNotificationsReq) (*pb.GetNotificationsResponse, error)
	GetAndMarkNotificationAsRead(context.Context, *pb.GetAndMarkNotificationAsReadReq) (*pb.GetAndMarkNotificationAsReadRes, error)
}
