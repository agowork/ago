package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/ago/apps/api/bootstrap"
	"github.com/rushairer/ago/config"
)

func main() {
	log.Println("starting...")

	server := gin.Default()
	bootstrap.SetupServer(server)

	log.Println("running...")
	if err := server.Run(fmt.Sprintf(":%s", config.ServerPort)); err != nil {
		log.Println(err)
	}
}
