package postgres

import (
	"auth-service/config"
	"auth-service/storage"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgresStorage struct {
	db *sql.DB
}

func ConnectDB() (storage.IStorage, error) {
	conf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Postgres.DB_HOST, conf.Postgres.DB_PORT, conf.Postgres.DB_USER, conf.Postgres.DB_NAME, conf.Postgres.DB_PASSWORD)
	log.Printf("connecting to postgres: %s\n", conDb)
	db, err := sql.Open("postgres", conDb)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &postgresStorage{db: db}, nil
}

func (p *postgresStorage) Close() {
	p.db.Close()
}

func (p *postgresStorage) User() storage.IUserStorage {
	return NewUserRepo(p.db)
}

func (p *postgresStorage) Notifications() storage.INotificationStorage {
	return NewNotificationsRepository(p.db)
}