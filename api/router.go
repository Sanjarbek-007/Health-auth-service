package api

import (
	_ "auth-service/api/docs"
	"auth-service/api/handler"
	"auth-service/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authorazation
// @version 1.0
// @description Authorazation API
// @host localhost:8081
// @BasePath /auth
func NewRouter(s storage.IStorage) *gin.Engine {
	h := handler.NewHandler(s)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/refresh", h.RefreshToken)
	auth.POST("/logout", h.Logout)

	auth.POST("/change-password", h.ChangePassword)
	auth.POST("/forgot-password", h.ForgotPassword)
	auth.POST("/reset-password", h.ResetPassword)

	return r
}
