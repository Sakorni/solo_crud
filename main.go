package main

import (
	"self_crud/server"
)

func main() {
	h := server.Handler{}
	server := h.InitHandler()
	server.Run(":80")
}
