package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"the-list/controller"
)

func SetupRoutes() *httprouter.Router {
	router := httprouter.New()

	// Static files
	router.ServeFiles("/static/*filepath", http.Dir("./web/public"))

	router.GET("/api/v1/shows", controller.GetShows)
	router.POST("/api/v1/shows", controller.PostShow)

	router.GET("/api/v1/search/shows", controller.SearchShow)

	// TODO, documentation is generated at: /doc/index.html
	// Is it possible to clean this up?
	router.GET("/doc/:any", controller.SwaggerHandler)

	router.GET("/", controller.Home)

	return router
}
