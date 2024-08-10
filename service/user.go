package service

import (
	pb "auth-service/genproto/user"
	"auth-service/logs"
	"auth-service/storage"
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
