package server

import (
	"github.com/sirupsen/logrus"
)

func (s *Server) GetUserState(userId int) string {
	state, err := s.service.User.State(userId)
	if err != nil {
		logrus.Error("error: server: GetUserState: " + err.Error())
	}

	return state
}
