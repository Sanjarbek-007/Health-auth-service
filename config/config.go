package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Postgres PostgresConfig
	Redis    RedisConfig
	Server   ServerConfig
}

type PostgresConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_NAME     string
	DB_PASSWORD string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

type ServerConfig struct {
	AUTH_ROUTER_PORT  string
	AUTH_SERVICE_PORT string
	ACCESS_TOKEN_KEY  string
	REFRESH_TOKEN_KEY string
	EMAIL             string
	APP_KEY           string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		Postgres: PostgresConfig{
			DB_HOST:     cast.ToString(coalesce("DB_HOST", "postgresdb")),
			DB_PORT:     cast.ToString(coalesce("DB_PORT", "5432")),
			DB_USER:     cast.ToString(coalesce("DB_USER", "macbookpro")),
			DB_NAME:     cast.ToString(coalesce("DB_NAME", "auth")),
			DB_PASSWORD: cast.ToString(coalesce("DB_PASSWORD", "1111")),
		},
		Redis: RedisConfig{
			Address:  cast.ToString(coalesce("REDIS_ADDRESS", "localhost:6379")),
			Password: cast.ToString(coalesce("REDIS_PASSWORD", "")),
			DB:       cast.ToInt(coalesce("REDIS_DB", "0")),
		},
		Server: ServerConfig{
			AUTH_ROUTER_PORT:  cast.ToString(coalesce("AUTH_ROUTER_PORT", ":8081")),
			AUTH_SERVICE_PORT: cast.ToString(coalesce("AUTH_SERVICE_PORT", ":50051")),
			ACCESS_TOKEN_KEY:  cast.ToString(coalesce("ACCESS_TOKEN_KEY", "key")),
			REFRESH_TOKEN_KEY: cast.ToString(coalesce("REFRESH_TOKEN_KEY", "key")),
			EMAIL:             cast.ToString(coalesce("EMAIL", "sanjarbekabdurahmonov123@gmail.com")),
			APP_KEY:           cast.ToString(coalesce("APP_KEY", "j e h d i a p p c b a b q n x p")),
		},
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
