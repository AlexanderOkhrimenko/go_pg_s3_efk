package main

import (
	"github.com/sirupsen/logrus"
	handler "go_pg_s3_efk/internal/handler"
	"go_pg_s3_efk/internal/repository"
	"go_pg_s3_efk/internal/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.NewPostgresDB()
	if err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes()

	logrus.Println("Server Start")
}
