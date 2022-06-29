package main

import (
	"fmt"
	"log"

	"forum/architecture/repository"
	"forum/architecture/service"
	"forum/architecture/web/handler"
	"forum/architecture/web/server"
	"forum/internal/database"
	"forum/internal/envfiller"

	_ "github.com/mattn/go-sqlite3"
)

const FILE_CONFIGS = "configs.env"

func main() {
	dbConf, servConf, handlerConf := GetConfigs()

	db, err := database.InitDatabase(dbConf)
	if err != nil {
		log.Fatalf("InitDatabase: %s\n", err)
	}

	repos := repository.NewRepo(db)
	services := service.NewService(repos)
	handlers, err := handler.NewMainHandler(services, handlerConf)
	if err != nil {
		log.Fatalln(err)
	}

	server := new(server.Server)
	if err := server.Run(servConf, handlers.InitRoutes(handlerConf)); err != nil {
		log.Fatalln(err)
	}
}

func GetConfigs() (*database.Configs, *server.Configs, *handler.Configs) {
	dbConf := new(database.Configs)
	err := envfiller.FillFieldsByEnvFile(FILE_CONFIGS, dbConf)
	if err != nil {
		log.Fatalf("GetConifgs from %q returns err: %v\n", FILE_CONFIGS, err)
	}

	servConf := new(server.Configs)
	err = envfiller.FillFieldsByEnvFile(FILE_CONFIGS, servConf)
	if err != nil {
		log.Fatalf("GetConifgs from %q returns err: %v\n", FILE_CONFIGS, err)
	}

	handlerConf := new(handler.Configs)
	err = envfiller.FillFieldsByEnvFile(FILE_CONFIGS, handlerConf)
	if err != nil {
		log.Fatalf("GetConifgs from %q returns err: %v\n", FILE_CONFIGS, err)
	}

	fmt.Printf("dbConf:         %+v\n", dbConf)
	fmt.Printf("servConf:       %+v\n", servConf)
	fmt.Printf("handlerConf:    %+v\n", handlerConf)
	return dbConf, servConf, handlerConf
}
