package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	User
	Photo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{NewUserRepository(db), NewPhotoRepository(db)}
}

type User interface {
	CreateUser(username string, userId int) error
	State(userId int) (string, error)
	ChangeState(userId int, newState string) (string, error)
	AddCallbackId(userId int, callbackId string) error
	AddCallbackData(callbackId, callbackData string) error
	GetCallbackData(userId int) (string, error)
	GetCallbackId(userId int) (int, error)
	DeleteCallback(userId int) error
}

type Photo interface {
	GetAllUserId() (*[]int, error)
	SetAlbumId(userId int, albumId string) error
	GetAlbumId(userId int) string
	DeleteAlbumId(albumId string) error
}
