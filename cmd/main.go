package main

// Замена стандартного пакета net/http:
// go get -u github.com/gin-gonic/gin

// Для работы с файлами конфигурации используется следующая библиотека, создает (configs/config.yml):
// go get -u github.com/spf13/viper

// Скачать библ-ку для миграций:
// https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
// brew install golang-migrate
// созд. папку schema с двумя файлами в ней:
// migrate create -ext sql -dir ./schema -seq init

// Для фикса грязных миграций:
//update schema_migrations set version='000001', dirty = false;

// сторонняя библиотека для работы с базой:
// go get github.com/jmoiron/sqlx

// Для получение паролей в приложениях используются переменные окружения.
// Библиотека для чтения файлов .env:
// https://github.com/joho/godotenv

// Библиотека для логирования, позволяет:
// - выводить логи в формате json
// - настраивать для логов свои кастомные поля
// - разделять логи по ур-ням
// https://github.com/sirupsen/logrus

// Для работы с jwt токенами:
// https://github.com/dgrijalva/jwt-go

// Swagger документация:
// https://github.com/swaggo/swag
// go get -u github.com/swaggo/swag/cmd/swag
// swag init -g cmd/main.go

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lukinairina90/REST_API_TODO_list"
	"github.com/lukinairina90/REST_API_TODO_list/pkg/handler"
	"github.com/lukinairina90/REST_API_TODO_list/pkg/repository"
	"github.com/lukinairina90/REST_API_TODO_list/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title REST_API_TODO_list App API
// @version 1.0
// @description API Server forTodoList Application

// @host localhost:8000
// BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(REST_API_TODO_list.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server %s", err.Error())
		}
	}()
	logrus.Println("Todo app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Todo Shutting Down")

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Fatalf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
