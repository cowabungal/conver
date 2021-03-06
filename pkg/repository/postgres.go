package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	usersTable = "users"
	callbacksTable = "callbacks"
	stateColumn = "state"
	callbackDataColumn = "callback_data"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(5 * time.Second)
	err = db.Ping()

	return db, err
}
