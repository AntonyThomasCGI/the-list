package web

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
  logger "github.com/sirupsen/logrus"
)


// getShows godoc
// @Summary      List all shows
// @Description  Get all shows currently stored in list
// @Tags         shows
// @Produce      json
// @Success      200  {list}  []Show
// @Router       /shows       [get]
func getShows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  shows, err := getItems()
  if err != nil {
    logger.Error(err)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(shows)
}


// postShow godoc
// @Summary      Add new show
// @Description  Add a new show to the list
// @Tags         shows
// @Accept       json
// @Success      200
// @Router       /shows [post]
func postShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  decoder := json.NewDecoder(r.Body)
  var t Show
  err := decoder.Decode(&t)
  if err != nil {
    logger.Error("Could not read body of request")
    return
  }
  logger.Info("Got data:")
  logger.Info(fmt.Sprintf("ID: %s", t.ID))
  logger.Info(fmt.Sprintf("Title: %s", t.Title))
  logger.Info(fmt.Sprintf("Author: %s", t.Author))
}

