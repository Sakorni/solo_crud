package main

import (
	"os"
	"self_crud/repository"
	"self_crud/server"
	"self_crud/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error occured during parsing .env file: %s", err.Error())
	}
	dbConfig := repository.DBConfig{
		User:     "root",
		Password: os.Getenv("MYSQL_ROOT_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
		Address:  "localhost:3308",
	}
	db, err := repository.CreateConnection(dbConfig)
	if err != nil {
		logrus.Fatalf("Error occured during setting db connection: %s", err.Error())
	}
	rep := repository.NewRepository(db)
	service := service.NewService(rep, &service.JWTService{})
	h := server.NewHandler(service)
	server := h.InitHandler()
	server.Run(":" + os.Getenv("APP_PORT"))
}
