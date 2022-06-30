package main

import (
	"log"
	"os"

	"forum/architecture/repository"
	"forum/architecture/service"
	"forum/architecture/web/handler"
	"forum/architecture/web/server"
)

func main() {
	repos := repository.NewRepo(nil)
	services := service.NewService(repos)
	handlers, err := handler.NewMainHandler(services)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	server := new(server.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
