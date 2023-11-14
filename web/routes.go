package web

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/julienschmidt/httprouter"
  logger "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
  Message string `json:"message"`
}


// getShows godoc
// @Summary      List all shows
// @Description  Get all shows currently stored in list
// @Tags         shows
// @Produce      json
// @Success      200  {List}  []Show
// @Failure      400  {object} ErrorResponse
// @Router       /shows       [get]
func getShows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  w.Header().Set("Content-Type", "application/json")

  shows, err := getItems()
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
func postShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  w.Header().Set("Content-Type", "application/json")
  var show Show
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

  id, err := saveItem(show)
  if err != nil {
    logger.Error("Could not read body of request: ", err)

    w.WriteHeader(http.StatusInternalServerError)

    resp := ErrorResponse{Message: err.Error()}
    json.NewEncoder(w).Encode(resp)

    return
  }
  show.ID = id
  json.NewEncoder(w).Encode(show)
}

