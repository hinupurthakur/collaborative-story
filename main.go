package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hinupurthakur/collaborative-story/api"
	"github.com/hinupurthakur/collaborative-story/db"
	"github.com/hinupurthakur/collaborative-story/logging"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warnln(".env not found, falling back to OS variables", err)
	}
	logging.Logger()
	db.InitDB()
}

func main() {
	port := os.Getenv("SERVER_PORT")
	r := api.CreateRoutes()
	log.Infof("starting server listening on :%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Errorln("ListenAndServe Errors:", err)
	}
}
