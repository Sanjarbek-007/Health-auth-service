package service

import (
	pb "auth-service/genproto/user"
	"auth-service/pkg/logger"
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
		logger:  logger.NewLogger(),
	}
}
