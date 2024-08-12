package storage

import (
	pb "auth-service/genproto/user"
	"auth-service/models"
	"context"
)

type IStorage interface {
	User() IUserStorage
	Close()
}

type IUserStorage interface {
	CreateUser(context.Context, *models.RegisterReq) (*models.RegisterResp, error)
	GetUserDetails(context.Context, string) (*models.UserDetails, error)
	GetUserProfile(context.Context, *pb.GetProfileReq) (*pb.GetProfileRes, error)
	UpdateProfile(context.Context, *pb.UpdateProfileReq) (*pb.Void, error)
	UpdatePassword(context.Context, *pb.UpdatePasswordReq) error
	GetUSerByEmail(context.Context, *pb.GetUSerByEmailReq) (*pb.FilterUsers, error)
	ChangePassword(ctx context.Context, userID, hashedPassword string) error
	GetByUserID(ctx context.Context, userID string) (*pb.User, error)
}
