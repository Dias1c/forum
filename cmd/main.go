package main

import (
	"fmt"
	"log"

	"forum/architecture/repository"
	"forum/architecture/service"
	"forum/architecture/web/handler"
	"forum/architecture/web/server"
	"forum/internal/envfiller"
)

const FILE_CONFIGS = "config.env"

func main() {
	servConf, handlerConf := GetConfigs()
	repos := repository.NewRepo(nil)
	services := service.NewService(repos)
	handlers, err := handler.NewMainHandler(services, handlerConf)
	if err != nil {
		log.Fatal(err)
	}

	server := new(server.Server)
	if err := server.Run(servConf, handlers.InitRoutes(handlerConf)); err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}
}

func GetConfigs() (*server.Configs, *handler.Configs) {
	servConf := new(server.Configs)
	err := envfiller.FillFieldsByEnvFile(FILE_CONFIGS, servConf)
	if err != nil {
		log.Fatalf("GetConifgs from %q returns err: %v\n", FILE_CONFIGS, err)
	}

	handlerConf := new(handler.Configs)
	err = envfiller.FillFieldsByEnvFile(FILE_CONFIGS, handlerConf)
	if err != nil {
		log.Fatalf("GetConifgs from %q returns err: %v\n", FILE_CONFIGS, err)
	}
	fmt.Printf("servConf: %+v\n", servConf)
	fmt.Printf("handlerConf: %+v\n", handlerConf)
	return servConf, handlerConf
}
