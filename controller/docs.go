package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func SwaggerHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	httpSwagger.WrapHandler(res, req)
}
