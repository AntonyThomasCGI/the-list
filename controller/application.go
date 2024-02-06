package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	logger "github.com/sirupsen/logrus"

	"the-list/db"
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

// getShows godoc
// @Summary      List all shows
// @Description  Get all shows currently stored in list
// @Tags         shows
// @Produce      json
// @Success      200  {List}  []Show
// @Failure      400  {object} ErrorResponse
// @Router       /shows       [get]
func GetShows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	shows, err := db.GetItems()
	if err != nil {
		logger.Error(err)

		w.WriteHeader(http.StatusInternalServerError)

		resp := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(resp)

		return
	}
	json.NewEncoder(w).Encode(shows)
}

// postShow godoc
// @Summary      Add new show
// @Description  Add a new show to the list
// @Tags         shows
// @Accept       json
// @Success      200 {object} Show
// @Failure      400 {object} ErrorResponse
// @Router       /shows [post]
func PostShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var show db.Show
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&show)
	if err != nil {
		logger.Error("Could not read body of request: ", err)

		w.WriteHeader(http.StatusBadRequest)

		resp := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(resp)

		return
	}
	logger.Info("Got data:")
	logger.Info(fmt.Sprintf("Title: %s", show.Title))
	logger.Info(fmt.Sprintf("Author: %s", show.Author))

	id, err := db.SaveItem(show)
	if err != nil {
		logger.Error("Error writing entry to db: ", err)

		w.WriteHeader(http.StatusInternalServerError)

		resp := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(resp)

		return
	}
	show.ID = id
	json.NewEncoder(w).Encode(show)
}
