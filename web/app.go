package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	logger "github.com/sirupsen/logrus"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "the-list/docs"
)

// @title           The List API
// @version         1.0
// @description     API for curating a movie and TV show watch list.

// @contact.name   Antony Thomas

// @host      localhost:8000
// @BasePath  /api/v1
func Start() {
	router := httprouter.New()

	// Static files
	router.ServeFiles("/static/*filepath", http.Dir("./public"))

	router.GET("/api/v1/shows", getShows)
	router.POST("/api/v1/shows", postShow)

	// TODO, documentation is generated at: /doc/index.html
	// Is it possible to clean this up?
	router.GET("/doc/:any", swaggerHandler)

	router.GET("/", home)

	logger.Info("Listen on 8000")
	logger.Fatal(http.ListenAndServe(":8000", router))
}

func swaggerHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
