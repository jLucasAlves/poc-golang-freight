package main

import (
	"github.com/dafiti-group/template-golang/configs"
)

func main() {
	server := configs.GetServer()

	server.Run()
}
