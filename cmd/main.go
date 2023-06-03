package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"task-management/internal/handler"
	"task-management/internal/repository"
	"task-management/internal/repository/database/mysql"
	"task-management/internal/server"
	"task-management/internal/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error occurred while initializating config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occurred while loading env variables: %s", err.Error())
	}

	db, err := mysql.NewMySQLDb(mysql.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("error occured while initializating database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Task Management started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Task Management Shutting Down")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Fatalf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("error occurred on database closing: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
