package service

import (
	"conver/pkg/repository"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username string, userId int) error {
	err := CreateDir(userId)
	if err != nil {
		logrus.Error("UserService: CreateUser: CreateDir: " + err.Error())
	}
	return s.repo.User.CreateUser(username, userId)
}

func (s *UserService) State(userId int) (string, error) {
	return s.repo.User.State(userId)
}

func (s *UserService) ChangeState(userId int, newState string) (string, error) {
	return s.repo.User.ChangeState(userId, newState)
}

func (s *UserService) AddCallbackId(userId int, callbackId string) error {
	return s.repo.User.AddCallbackId(userId, callbackId)
}

func (s *UserService) AddCallbackData(callbackId, callbackData string) error {
	return s.repo.User.AddCallbackData(callbackId, callbackData)
}

func (s *UserService) GetCallbackData(userId int) (string, error) {
	return s.repo.User.GetCallbackData(userId)
}

func (s *UserService) GetCallbackId(userId int) (int, error) {
	return s.repo.User.GetCallbackId(userId)
}

func (s *UserService) DeleteCallback(userId int) error {
	return s.repo.User.DeleteCallback(userId)
}
