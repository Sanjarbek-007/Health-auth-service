package handler

import (
	"auth-service/api/email"
	"auth-service/api/tokens"
	pb "auth-service/genproto/user"
	"auth-service/models"
	rdb "auth-service/storage/redis"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Registers user
// @Description Registers a new user
// @Tags auth
// @Param user body models.RegisterReq true "User data"
// @Success 200 {object} models.RegisterResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	h.Log.Info("Register handler is invoked")

	var req models.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		handlerError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	passByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		handlerError(c, h, err, "error hashing password", http.StatusInternalServerError)
		return
	}
	req.Password = string(passByte)

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	resp, err := h.Storage.User().CreateUser(ctx, &req)
	if err != nil {
		handlerError(c, h, err, "failed to register user", http.StatusInternalServerError)
		return
	}

	h.Log.Info("Register handler is completed successfully")
	c.JSON(http.StatusOK, resp)
}

// Login godoc
// @Summary Login user
// @Description Logs user in
// @Tags auth
// @Param user body models.LoginReq true "User data"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	h.Log.Info("Login handler is invoked")

	var req models.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		handlerError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	resp, err := h.Storage.User().GetUserDetails(ctx, req.Email)
	if err != nil {
		handlerError(c, h, err, "error fetching user", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(req.Password)); err != nil {
		handlerError(c, h, err, "invalid credentials", http.StatusBadRequest)
		return
	}

	accessToken, err := tokens.GenerateAccessToken(resp.Id, resp.Role)
	if err != nil {
		handlerError(c, h, err, "error generating access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := tokens.GenerateRefreshToken(resp.Id)
	if err != nil {
		handlerError(c, h, err, "error generating refresh token", http.StatusInternalServerError)
		return
	}

	h.Log.Info("Login handler is completed successfully")
	c.JSON(http.StatusOK, models.Tokens{
		AccessToken:  accessToken,
		ExpiresAt:    time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		RefreshToken: refreshToken,
	})
}

// RefreshToken godoc
// @Summary Refreshes token
// @Description Refreshes refresh token
// @Tags auth
// @Param data body models.RefreshToken true "Refresh token"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /refresh [post]
func (h *Handler) RefreshToken(c *gin.Context) {
	h.Log.Info("Refresh token handler is invoked")

	var req models.RefreshToken
	if err := c.ShouldBind(&req); err != nil {
		handlerError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	valid, err := tokens.ValidateRefreshToken(req.RefreshToken)
	if !valid || err != nil {
		handlerError(c, h, err, "error validating refresh token", http.StatusInternalServerError)
		return
	}

	id, err := tokens.ExtractRefreshUserID(req.RefreshToken)
	if err != nil {
		handlerError(c, h, err, "error extracting user id from refresh token", http.StatusInternalServerError)
		return
	}

	u, err := h.Storage.User().GetUserProfile(ctx, &pb.GetProfileReq{UserId: id})
	if err != nil {
		handlerError(c, h, err, "error fetching user", http.StatusInternalServerError)
		return
	}

	accessToken, err := tokens.GenerateAccessToken(u.UserId, u.Role)
	if err != nil {
		handlerError(c, h, err, "error generating access token", http.StatusInternalServerError)
		return
	}

	h.Log.Info("Refresh token handler is completed successfully")
	c.JSON(http.StatusOK, models.Tokens{
		AccessToken:  accessToken,
		ExpiresAt:    time.Now().Add(1 * time.Hour).Format(time.RFC3339),
		RefreshToken: req.RefreshToken,
	})
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Sends password reset code to user's email
// @Tags auth
// @Param data body models.ForgotPasswordReq true "User email"
// @Success 200 {string} string "Password reset code sent successfully"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /forgot-password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	h.Log.Info("Forgot password handler is invoked")

	var req models.ForgotPasswordReq
	if err := c.ShouldBind(&req); err != nil {
		handlerError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	_, err := email.SendEmail(req.Email)
	if err != nil {
		handlerError(c, h, err, "error sending email", http.StatusInternalServerError)
		return
	}

	h.Log.Info("Forgot password handler is completed successfully")
	c.JSON(http.StatusOK, "Password reset instructions sent to your email")
}

// ResetPassword godoc
// @Summary Reset password
// @Description Resets user password
// @Tags auth
// @Param data body models.ResetPasswordReq true "Reset password data"
// @Success 200 {string} string "Password reset successfully"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /reset-password [post]
// func (h *Handler) ResetPassword(c *gin.Context) {}
func (h *Handler) ResetPassword(c *gin.Context) {
	h.Log.Info("Reset password handler is invoked")

	var req models.ResetPasswordReq
	if err := c.ShouldBind(&req); err != nil {
		handlerError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	code, err := rdb.GetCode(ctx, req.Email)
	if err != nil {
		handlerError(c, h, err, "error getting code", http.StatusInternalServerError)
		return
	}

	if code != req.Code {
		handlerError(c, h, err, "invalid code", http.StatusBadRequest)
		return
	}

	err = rdb.DeleteCode(ctx, req.Email)
	if err != nil {
		handlerError(c, h, err, "error deleting code", http.StatusInternalServerError)
		return
	}

	u, err := h.Storage.User().GetUserDetails(ctx, req.Email)
	if err != nil {
		handlerError(c, h, err, "user not found", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		handlerError(c, h, err, "error hashing password", http.StatusInternalServerError)
		return
	}

	err = h.Storage.User().UpdatePassword(ctx, &pb.UpdatePasswordReq{
		Id:       u.Id,
		Password: string(hashedPassword),
	})
	if err != nil {
		handlerError(c, h, err, "error updating password", http.StatusInternalServerError)
		return
	}

	h.Log.Info("Reset password handler is completed successfully")
	c.JSON(http.StatusOK, "Password reset successfully")
}

// ChangePassword godoc
// @Summary Changes password
// @Description Changes user password
// @Tags auth
// @Param data body models.ChangePasswordReq true "Change password data"
// @Success 200 {string} string "Password changed successfully"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /change-password [post]
func (h *Handler) ChangePassword(c *gin.Context) {
	h.Log.Info("Change password handler is invoked")

	var req models.ChangePasswordReq
	if err := c.ShouldBind(&req); err != nil {
		handlerError(c, h, err, "invalid data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	u, err := h.Storage.User().GetUserDetails(ctx, req.Email)
	if err != nil {
		handlerError(c, h, err, "user not found", http.StatusBadRequest)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.CurrentPassword)); err != nil {
		handlerError(c, h, err, "incorrect password", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		handlerError(c, h, err, "error hashing password", http.StatusInternalServerError)
		return
	}

	err = h.Storage.User().ChangePassword(ctx, u.Id, string(hashedPassword))
	if err != nil {
		handlerError(c, h, err, "error changing password", http.StatusInternalServerError)
		return
	}

	h.Log.Info("Change password handler is completed successfully")
	c.JSON(http.StatusOK, "Password changed successfully")
}

// Logout godoc
// @Summary Logouts user
// @Description Logouts user
// @Tags auth
// @Param email query string true "User email"
// @Success 200 {string} string "User logged out successfully"
// @Failure 400 {object} string "Invalid user id"
// @Failure 500 {object} string "Server error while processing request"
// @Router /logout [post]
func (h *Handler) Logout(c *gin.Context) {
	h.Log.Info("Logout handler is invoked")

	email := c.Query("email")
	if email == "" {
		handlerError(c, h, errors.New("no email provided"), "invalid email", http.StatusBadRequest)
		return
	}

	h.Log.Info("User logged out successfully")
	c.JSON(http.StatusOK, "User logged out successfully")
}
