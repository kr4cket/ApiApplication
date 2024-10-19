package app

import (
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"gitlab.dev-masters.tech/shepelev-projects/sisso/sisso-api/internal/handlers"
	"gitlab.dev-masters.tech/shepelev-projects/sisso/sisso-api/internal/server"
	"log"
	"os"
)

func Run() {
	handler := handlers.New()
	serv := new(server.Server)

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("error loading env vars: %s", err.Error())
	}

	if err := serv.Run(os.Getenv("API_PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running Http-Server: %s", err.Error())
	}
}
