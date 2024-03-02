package main

import (
	"log"
	"os"

	"github.com/exxxception/restful-service/internal/config"
	"github.com/exxxception/restful-service/internal/handlers"
	"github.com/exxxception/restful-service/internal/http_server"
	"github.com/exxxception/restful-service/internal/services"
	"github.com/exxxception/restful-service/pkg/repository"
)

//	@title			REST API Service
//	@version		1.0
//	@description	Swagger API for REST API Service.

//	@host		localhost:8080
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	cfg := config.LoadConfig()

	repos := repository.NewRepository(cfg.DB.DSN)
	services := services.NewServices(repos)
	handlers := handlers.NewHandlers(services, cfg.ApiKey)

	srv := new(http_server.Server)
	if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
		os.Exit(1)
	}
}
