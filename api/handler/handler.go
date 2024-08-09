package handler

import (
	"auth-service/pkg/logger"
	"auth-service/storage"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	Storage storage.IStorage
	Log     *slog.Logger
}

func NewHandler(s storage.IStorage) *Handler {
	return &Handler{
		Storage: s,
		Log:     logger.NewLogger(),
	}
}

func handlerError(c *gin.Context, h *Handler, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})
	h.Log.Error(er)
}
