package postgres

import (
	pb "auth-service/genproto/user"
	"auth-service/models"
	"auth-service/pkg/logger"
	"auth-service/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type UserRepo struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewUserRepo(DB *sql.DB) storage.IUserStorage {
	return &UserRepo{DB: DB, Log: logger.NewLogger()}
}

func (u *UserRepo) CreateUser(ctx context.Context, req *models.RegisterReq) (*models.RegisterResp, error) {
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var userID, createdAtStr string
	userQuery := `INSERT INTO users (username, email, password_hash, full_name, native_language, role) 
                  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`
	err = tx.QueryRowContext(ctx, userQuery, req.Username, req.Email, req.Password, req.Fullname, req.NativeLanguage, req.Role).Scan(&userID, &createdAtStr)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	createdAt, err := time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created_at: %w", err)
	}

	return &models.RegisterResp{
		Id:             userID,
		Username:       req.Username,
		Email:          req.Email,
		Fullname:       req.Fullname,
		NativeLanguage: req.NativeLanguage,
		CreatedAt:      createdAt.String(),
	}, nil
}

func (u *UserRepo) GetUserDetails(ctx context.Context, email string) (*models.UserDetails, error) {
	query := `SELECT id, username, password_hash, role FROM users WHERE email = $1`

	var id, username, passwordHash, role string

	err := u.DB.QueryRowContext(ctx, query, email).Scan(&id, &username, &passwordHash, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &models.UserDetails{
		Id:       id,
		Username: username,
		Password: passwordHash,
		Role:     role,
	}, nil

}

func (u *UserRepo) GetUserProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileRes, error) {
	query := `SELECT id, username, email, full_name, native_language, role FROM users WHERE id = $1`
	var id, username, email, fullName, nativeLanguage, role string

	err := u.DB.QueryRowContext(ctx, query, req.Id).Scan(&id, &username, &email, &fullName, &nativeLanguage, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &pb.GetProfileRes{
		Id:             id,
		Username:       username,
		Email:          email,
		Fullname:       fullName,
		NativeLanguage: nativeLanguage,
		Role:           role,
	}, nil
}

func (u *UserRepo) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.Void, error) {
	query := `UPDATE users SET full_name = $1, native_language = $2 WHERE id = $3`
	_, err := u.DB.ExecContext(ctx, query, req.Fullname, req.NativeLanguage, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (u *UserRepo) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq) error {
	query := `update users set password_hash=$1 where id=$2 and deleted_at=0`
	result, err := u.DB.ExecContext(ctx, query, req.Password, req.Id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (u *UserRepo) GetUSerByEmail(ctx context.Context, req *pb.GetUSerByEmailReq) (*pb.FilterUsers, error) {
	query := `SELECT id, username, email, full_name, native_language, role FROM users WHERE email=$1`
	var id, username, email, fullName, nativeLanguage, role string

	err := u.DB.QueryRowContext(ctx, query, req.Email).Scan(&id, &username, &email, &fullName, &nativeLanguage, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &pb.FilterUsers{
		Id:             id,
		Username:       username,
		Email:          email,
		Fullname:       fullName,
		NativeLanguage: nativeLanguage,
		Role:           role,
	}, nil
}

func (u *UserRepo) ChangePassword(ctx context.Context, userID, hashedPassword string) error {
	query := `update users set password_hash=$1 where id=$2 and deleted_at=0`
	_, err := u.DB.ExecContext(ctx, query, hashedPassword, userID)
	if err != nil {
		return err
	}

	return nil
}
