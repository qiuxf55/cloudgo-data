package main

import (
	"MySql/service"
)

func main() {
	server := service.NewServer()
	server.Run()
}