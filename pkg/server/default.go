package server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func (s *Server) text (m *telebot.Message) {
	logrus.Printf("Text from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	state := s.GetUserState(m.Sender.ID)
	switch state {
	case "default":
		s.bot.Send(m.Sender, "Нажми на кнопку 'Конвертация', чтобы приступить к загрузке фото.")
	}
}

/*func (s *Server) photo(m *telebot.Message)  {
	//wg := new(sync.WaitGroup)
	//rl := s.rl
	//rll := *(rl)
	//_ = rll.Take()
	s.photoGet(m)

	//nrl := ratelimit.New(100000)
	//s.rl = &nrl
	//wg.Wait()
}*/