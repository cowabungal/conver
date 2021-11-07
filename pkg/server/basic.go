package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"os"
)

func (s *Server) start(m *telebot.Message) {
	logrus.Printf("/start from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	err := s.service.User.CreateUser(m.Sender.Username, m.Sender.ID)
	if err != nil {
		logrus.Error("start: CreateUser: " + err.Error())
	}

	main := s.button.Main()
	s.bot.Send(m.Sender,"Привет. Я - бот, который сконвертирует все твои фотки в pdf.", &main)
}


func (s *Server) help(m *telebot.Message) {
	logrus.Printf("Помощь from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	s.service.User.ChangeState(m.Sender.ID, "default")
	s.bot.Send(m.Sender, "Слова поддержки). Если что-то пошло не так, или есть пожелания по улучшению бота, смело пишите @Cowabunga_a.")
}

func (s *Server) convert(m *telebot.Message) {
	logrus.Printf("Конвертация from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)
	s.bot.Send(m.Sender, "Отправь фотки одним сообщением.")
	s.service.ChangeState(m.Sender.ID, "convert")
}

func (s *Server) photo(m *telebot.Message) {
	logrus.Printf("Фото: from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	photo, err := s.bot.GetFile(&m.Photo.File)
	if err != nil {
		logrus.Error("photo: GetFile: " + err.Error())
	}
	n, err := s.service.Photo.Save(m.Sender.ID, photo)
	if err != nil {
		logrus.Error("photo: Save: " + err.Error())
	}

	s.bot.Send(m.Sender, fmt.Sprintf("Загружено %dшт фото. Нажми /convert, чтобы отправить на конвертацию.", n))
	s.service.ChangeState(m.Sender.ID, "photo")
}

func (s *Server) pdf(m *telebot.Message) {
	logrus.Printf("Pdf from: %s; id: %d; ms: %s", m.Sender.Username, m.Sender.ID, m.Text)

	err := s.service.Photo.GetPdf(m.Sender.ID)
	if err != nil {
		logrus.Errorf("pdf: GetPdf: " + err.Error())
	}

	_, err = os.Stat(fmt.Sprintf("./assets/%d/conver.pdf", m.Sender.ID))
	if !os.IsNotExist(err) {
		_, err = s.bot.Send(m.Sender, &telebot.Document{
			File:      telebot.File{FileLocal: fmt.Sprintf("./assets/%d/conver.pdf", m.Sender.ID)},
			Thumbnail: nil,
			Caption:   "📋 Твой pdf файл готов.\n" +
				"💬 Если есть пожелания по улучшению бота, пишите @Cowabunga_a.\n" +
				"✓ Официальный канал бота @botconver",
			MIME:      "",
			FileName:  "conver.pdf",
		})
		if err != nil {
			logrus.Error(err.Error())
		}
	}

	clearDirectory(fmt.Sprintf("./assets/%d/", m.Sender.ID))
}

func clearDirectory(directory string)  {
	readDirectory, _ := os.Open(directory)
	allFiles, _ := readDirectory.Readdir(0)

	for f := range(allFiles) {
		file := allFiles[f]

		fileName := file.Name()
		filePath := directory + fileName

		os.Remove(filePath)
	}
}

	//love u, детка
