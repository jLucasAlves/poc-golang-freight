package main

import (
	"dafiti-group/microservice/configs"
)

func main() {
	server := configs.GetServer()

	server.Run()
}
