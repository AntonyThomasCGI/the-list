package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/julienschmidt/httprouter"
	logger "github.com/sirupsen/logrus"

	"the-list/db"
)

// GetShows godoc
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

// PostShow godoc
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

// SearchShow godoc
func SearchShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	baseUrl := "https://api.themoviedb.org/3/search/multi"
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		msg := "TMDB_API_KEY not set!"
		logger.Error(msg)
		w.WriteHeader((http.StatusInternalServerError))
		resp := ErrorResponse{Message: msg}
		json.NewEncoder(w).Encode(resp)

		return
	}

	queryValues := r.URL.Query()

	logger.Info("==================================================")
	param, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(string(param))
	logger.Info(ps.ByName("query"))
	logger.Info(queryValues.Get("query"))

	endpoint := baseUrl + "?api_key=" + apiKey + "&query=" + url.QueryEscape(queryValues.Get("query"))
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		logger.Error(err)
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)

		w.WriteHeader(http.StatusInternalServerError)

		resp := ErrorResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(resp)

		return
	}

	defer res.Body.Close()

	//data, err := io.ReadAll(res.Body)
	//if err != nil {
	//	logger.Error(err)
	//	// TODO
	//	return
	//}
	//logger.Info(string(data))

	type rawResult struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Name        string `json:"name"`
		MediaType   string `json:"media_type"`
		ReleaseDate string `json:"release_data"`
	}

	respJson := struct {
		Results []*rawResult `json:"results"`
	}{}
	decodeErr := json.NewDecoder(res.Body).Decode(&respJson)
	if decodeErr != nil {
		logger.Error(decodeErr)
		// TODO
		return
	}

	type formattedResult struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		ReleaseDate string `json:"release_data"`
	}

	transformResult := []formattedResult{}
	for i := range respJson.Results {
		// Filter for media types in movie/tv show.
		validMediaType := false
		for _, a := range []string{"movie", "tv"} {
			if respJson.Results[i].MediaType == a {
				validMediaType = true
				break
			}
		}
		if !validMediaType {
			continue
		}
		// Movies have title, tv shows have name json fields for some dumb reason.
		title := respJson.Results[i].Title
		if title == "" {
			title = respJson.Results[i].Name
		}
		transformResult = append(transformResult, formattedResult{
			ID:          respJson.Results[i].ID,
			Title:       title,
			ReleaseDate: respJson.Results[i].ReleaseDate,
		})
	}
	logger.Info("++json resp:")
	logger.Info(transformResult[0].Title)
	logger.Info(transformResult[1].Title)
	logger.Info(transformResult[2].Title)
	json.NewEncoder(w).Encode(transformResult)
}
