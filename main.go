package main

import (
	"self_crud/server"
	"self_crud/service"
)

func main() {
	service := service.NewService()
	h := server.NewHandler(service)
	server := h.InitHandler()
	server.Run(":80")
}
