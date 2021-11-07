package server

import (
	"conver/pkg/server/buttons"
	"conver/pkg/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

type Server struct {
	service    *service.Service
	bot    *telebot.Bot
	button *buttons.Buttons
	upd chan *telebot.Update
}

func (s *Server) InitRoutes() {
	go s.processPhoto(s.upd)

	s.bot.Handle("Конвертация", s.convert)
	s.bot.Handle("Помощь", s.help)
	s.bot.Handle("/start", s.start)
	s.bot.Handle("/convert", s.pdf)
	s.bot.Handle(telebot.OnText, s.text)
	//s.bot.Handle(telebot.OnPhoto, s.photo)
}

func (s *Server) processPhoto(updCh chan *telebot.Update)  {
	for {
		select {
		case v1 := <-updCh:
			s.photo(v1.Message)
			time.Sleep(500 *time.Millisecond)
		}
	}
}

func (s *Server) Run() {
	s.InitRoutes()
	logrus.Info("The BotServer has successfully run")
	s.bot.Start()
}

func NewBotServer(s *service.Service) *Server {
	p := &telebot.LongPoller{Timeout: 15 * time.Second}
	updCh := make(chan *telebot.Update, 1000)
	poller := telebot.NewMiddlewarePoller(p, func(upd *telebot.Update) bool {
		if upd.Message != nil || upd.Callback != nil {
			if upd.Message.Photo != nil {
				updCh <- upd
				return false
			}
			return true
		}

		return false
	})

	b, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: poller,
	})

	if err != nil {
		log.Fatal(err)
		return nil
	}

	menu := telebot.ReplyMarkup{ResizeReplyKeyboard: true}
	bu := buttons.NewButtons(menu)
	return &Server{service: s, bot: b, button: bu, upd: updCh}
}
