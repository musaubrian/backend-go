package main

import (
	"github.com/musaubrian/backend-go/models"
	"github.com/musaubrian/backend-go/server"
)

func main() {
	models.Setup()
	server.SetupAndListen()
}
