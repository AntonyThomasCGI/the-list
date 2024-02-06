package main

import (
	"net/http"

	logger "github.com/sirupsen/logrus"

	_ "the-list/docs"
	"the-list/routers"
)

// @title           The List API
// @version         1.0
// @description     API for curating a movie and TV show watch list.

// @contact.name   Antony Thomas

// @host      localhost:8000
// @BasePath  /api/v1
func main() {
	router := routers.SetupRoutes()

	logger.Info("Listen on 8000")
	logger.Fatal(http.ListenAndServe(":8000", router))
}
