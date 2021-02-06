package main

import (
	"github.com/dafiti-group/golang-template-project/configs"
)

func main() {
	server := configs.GetServer()

	server.Run()
}
