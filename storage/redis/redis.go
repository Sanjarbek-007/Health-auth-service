package redis

import (
	"auth-service/config"
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func ConnectDB() *redis.Client {
	cfg := config.Load().Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return rdb
}

func StoreCode(ctx context.Context, email, code string) error {
	fmt.Println("Storing code for email: " + email + " with code: " + code)
	rdb := ConnectDB()

	err := rdb.Set(ctx, "lingualeap:"+email, code, time.Minute*3).Err()
	if err != nil {
		return errors.Wrap(err, "failed to store code in redis")

	}

	return nil
}

func GetCode(ctx context.Context, email string) (string, error) {
	rdb := ConnectDB()

	code, err := rdb.Get(ctx, "lingualeap:"+email).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New("code not found for " + email)
		}
		return "", errors.Wrap(err, "failed to get code from redis")
	}

	return code, nil
}

func DeleteCode(ctx context.Context, email string) error {
	rdb := ConnectDB()

	err := rdb.Del(ctx, "lingualeap:"+email).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete code from redis")
	}

	return nil
}
