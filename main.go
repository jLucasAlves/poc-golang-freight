package main

import (
	"github.com/dafiti-group/${{values.component_id}}/configs"
)

func main() {
	server := configs.GetServer()

	server.Run()
}
