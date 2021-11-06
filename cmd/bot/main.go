package main

import (
	"conver/pkg/repository"
	"conver/pkg/server"
	"conver/pkg/service"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// инициализация бд
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	srv := server.NewBotServer(services)

	//инициализация всех рабочих папок с userId
	CreateDirs(repository)

	//запуск бота
	srv.Run()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func CreateDirs(r *repository.Repository)  {
	usersId, err := r.Photo.GetAllUserId()
	if err != nil {
		logrus.Fatalf("failed to initialize dirs %s", err.Error())
	}

	for _, v := range *usersId {
		err = os.Mkdir(fmt.Sprintf("./assets/%d", v), 0o777)
		if err != nil {
			logrus.Error(err.Error())
		}
	}
}
