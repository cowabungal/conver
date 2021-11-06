package service

import (
	"conver/pkg/repository"
	"io"
)

type Service struct {
	User
	Photo
}

func NewService(repo *repository.Repository) *Service {
	return &Service{NewUserService(repo), NewPhotoService(repo)}
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
	Save(userId int, reader io.Reader) (int, error)
	SetAlbumId(userId int, albumId string) error
	GetAlbumId(userId int) string
	DeleteAlbumId(albumId string) error
	GetPdf(userId int) error
}
