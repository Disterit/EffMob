package main

import (
	"EffMob/logger"
	"EffMob/pkg/handler"
	"EffMob/pkg/repositroy"
	"EffMob/pkg/service"
	"EffMob/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"

	_ "EffMob/docs"
)

// @title Effective Mobile API
// @version 1.0
// @description Api server for test example

// @host localhost:8080
// @BasePath /

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Log.Error("Error loading .env file")
		return
	}

	if err := readInConfig(); err != nil {
		logger.Log.Error("Error loading config file file")
	}

	db := repositroy.Connection(repositroy.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	repo := repositroy.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(server.Api)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Log.Error("error to start server", err.Error())
	}

}

func readInConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
