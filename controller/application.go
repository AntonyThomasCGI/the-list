package controller

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	logger "github.com/sirupsen/logrus"
)

func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contents, err := os.ReadFile("web/public/index.html")
	if err != nil {
		logger.Error(err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(contents)
}
