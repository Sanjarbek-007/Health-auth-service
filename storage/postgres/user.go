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

func (u *UserRepo) CreateUser(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterRes, error) {
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	var userID, createdAtStr string
	userQuery := `INSERT INTO users (first_name, last_name, email, password_hash, date_of_birth, gender, role) 
                  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at`
	err = tx.QueryRowContext(ctx, userQuery, req.FirstName, req.LastName, req.Email, req.Password, req.DateOfBirth, req.Gender, req.Role).Scan(&userID, &createdAtStr)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	_, err = time.Parse(time.RFC3339, createdAtStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created_at: %w", err)
	}

	return &pb.RegisterRes{
		UserId: userID,
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
	fmt.Println(userID, "/////////////////")
	query := `SELECT id, first_name, last_name, email, password_hash, role, date_of_birth, gender FROM users WHERE id=$1 AND deleted_at = 0`
	var id, firstName, lastName, email, passwordHash, role, dateOfBirth, gender string

	err := u.DB.QueryRowContext(ctx, query, userID).Scan(&id, &firstName, &lastName, &email, &passwordHash, &role, &dateOfBirth, &gender)
	if err != nil {
		fmt.Println(1)
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	fmt.Println(firstName)
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
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) CreateNotifications(ctx context.Context, req *pb.CreateNotificationsReq) (*pb.CreateNotificationsRes, error) {
	var id string
	fmt.Println(req.UserId, req.GetMessage(), false)
	err := u.DB.QueryRowContext(ctx, `
		INSERT INTO notifications (user_id, messages, is_read) 
		VALUES ($1, $2, $3) RETURNING id`,
		req.UserId, req.GetMessage(), false,
	).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert notification: %w", err)
	}

	return &pb.CreateNotificationsRes{Id: id}, nil
}

func (u *UserRepo) GetAllNotifications(ctx context.Context, req *pb.GetNotificationsReq) (*pb.GetNotificationsResponse, error) {
	rows, err := u.DB.QueryContext(ctx, `
		SELECT id, user_id, messages 
		FROM notifications 
		WHERE user_id = $1`, req.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("failed to query notifications: %w", err)
	}
	defer rows.Close()

	var notifications []*pb.Notification
	for rows.Next() {
		var notification pb.Notification
		if err := rows.Scan(&notification.Id, &notification.UserId, &notification.Message); err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}
		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &pb.GetNotificationsResponse{Notifications: notifications}, nil
}

func (u *UserRepo) GetAndMarkNotificationAsRead(ctx context.Context, req *pb.GetAndMarkNotificationAsReadReq) (*pb.GetAndMarkNotificationAsReadRes, error) {
	rows, err := u.DB.QueryContext(ctx, `
		SELECT id, user_id, messages 
		FROM notifications 
		WHERE user_id = $1 AND is_read = FALSE`, req.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("failed to query notifications: %w", err)
	}
	defer rows.Close()

	var notifications []*pb.Notification
	for rows.Next() {
		var notification pb.Notification
		if err := rows.Scan(&notification.Id, &notification.UserId, &notification.Message); err != nil {
			return nil, fmt.Errorf("failed to scan notification: %w", err)
		}
		notifications = append(notifications, &notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	for _, notif := range notifications {
		_, err := tx.ExecContext(ctx, `
			UPDATE notifications 
			SET is_read = TRUE 
			WHERE id = $1`, notif.GetId())
		if err != nil {
			return nil, fmt.Errorf("failed to update notification %s: %w", notif.GetId(), err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &pb.GetAndMarkNotificationAsReadRes{Notifications: notifications}, nil
}
