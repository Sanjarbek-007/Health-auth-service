package postgres

import (
	pb "auth-service/genproto/user"
	"auth-service/logs"
	"auth-service/models"
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
	return &UserRepo{DB: DB, Log: logs.NewLogger()}
}

func (u *UserRepo) CreateUser(ctx context.Context, req *models.RegisterReq) (*models.RegisterResp, error) {
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var userID, createdAtStr string
	userQuery := `INSERT INTO users (first_name, last_name, email, password_hash, date_of_birth, gender) 
                  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`
	err = tx.QueryRowContext(ctx, userQuery, req.FirstName, req.LastName, req.Email, req.Password, req.DateOfBirth, req.Gender).Scan(&userID, &createdAtStr)
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
		Id:          userID,
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Gender:      req.Gender,
		CreatedAt:   createdAt.String(),
	}, nil
}

func (u *UserRepo) GetUserDetails(ctx context.Context, email string) (*models.UserDetails, error) {
	query := `SELECT id, password_hash, role FROM users WHERE email = $1 AND deleted_at = 0`

	var id, passwordHash, role string

	err := u.DB.QueryRowContext(ctx, query, email).Scan(&id, &passwordHash, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &models.UserDetails{
		Id:       id,
		Email:    email,
		Password: passwordHash,
		Role:     role,
	}, nil

}

func (u *UserRepo) GetUserProfile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileRes, error) {
	query := `SELECT id, first_name, last_name, email, gender, role, date_of_birth FROM users WHERE id = $1 AND deleted_at = 0`
	var id, firstName, lastName, email, gender, role, dateOfBirth string

	err := u.DB.QueryRowContext(ctx, query, req.UserId).Scan(&id, &firstName, &lastName, &email, &gender, &role, &dateOfBirth)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &pb.GetProfileRes{
		UserId:      id,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		Gender:      gender,
		Role:        role,
		DateOfBirth: dateOfBirth,
	}, nil
}

func (u *UserRepo) UpdateProfile(ctx context.Context, req *pb.UpdateProfileReq) (*pb.Void, error) {
	var email, firstName, lastName, dateOfBirth, gender string
	get := `SELECT email, first_name, last_name, date_of_birth, gender FROM users WHERE id = $1 AND deleted_at = 0`
	err := u.DB.QueryRowContext(ctx, get, req.Id).Scan(&email, &firstName, &lastName, &dateOfBirth, &gender)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Email != "" {
		req.Email = email
	}
	if req.FirstName != "" {
		req.FirstName = firstName
	}
	if req.LastName != "" {
		req.LastName = lastName
	}
	if req.DateOfBirth != "" {
		req.DateOfBirth = dateOfBirth
	}
	if req.Gender != "" {
		req.Gender = gender
	}

	query := `UPDATE users SET email = $1, first_name = $2, last_name = $3, date_of_birth = $4, gender = $5 WHERE id = $6 AND deleted_at=0`
	_, err = u.DB.ExecContext(ctx, query, req.Email, req.FirstName, req.LastName, req.DateOfBirth, req.Gender, req.Id)
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
	query := `SELECT id, first_name, last_name, email, date_of_birth, role, gender FROM users WHERE email=$1 AND deleted_at = 0`
	var id, firstName, lastName, email, dateOfBirth, role, gender string

	err := u.DB.QueryRowContext(ctx, query, req.Email).Scan(&id, &firstName, &lastName, &email, &dateOfBirth, &role, &gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &pb.FilterUsers{
		UserId:      id,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
		Gender:      gender,
		Role:        role,
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

func (u *UserRepo) GetByUserID(ctx context.Context, userID string) (*pb.User, error) {
	query := `SELECT id, first_name, last_name, email, password_hash, role, date_of_birth, gender FROM users WHERE id=$1 AND deleted_at = 0`
	var id, firstName, lastName, email, passwordHash, role, dateOfBirth, gender string

	err := u.DB.QueryRowContext(ctx, query, userID).Scan(&id, &firstName, &lastName, &email, &passwordHash, &role, &dateOfBirth, &gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &pb.User{
		UserId:      id,
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
		Gender:      gender,
	}, err
}

func (u *UserRepo) DeleteUser(ctx context.Context, userID string) error {
	now := time.Now().Unix()
	query := `UPDATE users SET deleted_at = $1 WHERE id = $2 AND deleted_at = 0`
    _, err := u.DB.ExecContext(ctx, query, now, userID)
    if err!= nil {
        return err
    }

    return nil
}