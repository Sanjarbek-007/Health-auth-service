package handler

import (
	"auth-service/logs"
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
		Log:     logs.NewLogger(),
	}
}

func handlerError(c *gin.Context, h *Handler, err error, msg string, code int) {
	er := errors.Wrap(err, msg).Error()
	c.AbortWithStatusJSON(code, gin.H{"error": er})
	h.Log.Error(er)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
