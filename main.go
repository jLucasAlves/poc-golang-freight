package main

import (
	"github.com/dafiti-group/poc-golang-freight/configs"
)

func main() {
	server := configs.GetServer()

	server.Run()
}
