package main

import (
	"L2.18-calendar/internal/apiserver"
	"L2.18-calendar/internal/handler"
	"L2.18-calendar/internal/repository"
	"L2.18-calendar/internal/service"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	Init()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	//repo

	eventRepo := repository.NewRepository()
	eventService := service.NewService(eventRepo)
	eventHandler := handler.NewHandler(eventService)
	s := apiserver.New(config)
	s.ConfigureRouter(eventHandler)

	if err := s.Run(); err != nil {
		panic(err)
	}
}
