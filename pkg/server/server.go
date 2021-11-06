package server

import (
	"conver/pkg/server/buttons"
	"conver/pkg/service"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
	"gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
	"time"
)

type Server struct {
	service    *service.Service
	bot    *telebot.Bot
	button *buttons.Buttons
	rl *ratelimit.Limiter
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
			time.Sleep(100 *time.Millisecond)
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
	rl := ratelimit.New(100000) //per 100ms
	return &Server{service: s, bot: b, button: bu, rl: &rl, upd: updCh}
}
