package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"

	"github.com/julienschmidt/httprouter"

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
		writeErrorResponse(err.Error(), w)
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
		writeErrorResponse(err.Error(), w)
		return
	}

	id, err := db.SaveItem(show)
	if err != nil {
		writeErrorResponse(err.Error(), w)
		return
	}
	show.ID = id
	json.NewEncoder(w).Encode(show)
}

// TODO, doc string
func UpdateShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	if id == "" {
		writeErrorResponse("Missing required id parameter", w)
		return
	}

	body := map[string]interface{}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		msg := fmt.Sprintf("Error decoding json data: %s", err.Error())
		writeErrorResponse(msg, w)
		return
	}

	show := db.Show{}
	rt := reflect.TypeOf(show)
	rv := reflect.ValueOf(show)
	fields := []string{}
	for i := 0; i < rt.NumField(); i++ {
		fields = append(fields, rt.Field(i).Tag.Get("json"))
	}

	sanitizedData := map[string]interface{}{}
	for k, v := range body {
		// Check user field is a legit field on a show.
		exists := false
		fieldN := 0
		for i, field := range fields {
			if k == field {
				exists = true
				fieldN = i
				break
			}
		}
		if !exists {
			msg := fmt.Sprintf("Got unexpected field '%s'", k)
			writeErrorResponse(msg, w)
			return
		}
		dataType := rv.Field(fieldN).Type()
		castVal, ok := safeCast(v, dataType)
		if !ok {
			msg := fmt.Sprintf("Unexpected type for field '%s' expected: %s", k, dataType)
			writeErrorResponse(msg, w)
			return
		}
		sanitizedData[k] = castVal
	}

	if len(sanitizedData) == 0 {
		writeErrorResponse("No fields provided to update", w)
		return
	}

	err2 := db.UpdateItem(id, sanitizedData)
	if err2 != nil {
		msg := fmt.Sprintf("Failed to write to database: %s", err2.Error())
		writeErrorResponse(msg, w)
		return
	}
}

// SearchShow godoc
func SearchShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	baseUrl := "https://api.themoviedb.org/3/search/multi"
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		msg := "TMDB_API_KEY not set in environment!"
		writeErrorResponse(msg, w)
		return
	}
	queryValues := r.URL.Query()

	endpoint := baseUrl + "?api_key=" + apiKey + "&query=" + url.QueryEscape(queryValues.Get("query"))
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		msg := fmt.Sprintf("Error while setting up tmdb request: %s", err.Error())
		writeErrorResponse(msg, w)
		return
	}
	req.Header.Add("accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		msg := fmt.Sprintf("Failed to forward request to tmdb: %s", err.Error())
		writeErrorResponse(msg, w)
		return
	}

	defer res.Body.Close()

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
		msg := fmt.Sprintf("Error decoding tmdb response: %s", decodeErr.Error())
		writeErrorResponse(msg, w)
		return
	}

	type formattedResult struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		ReleaseDate string `json:"release_data"`
	}

	transformResult := []formattedResult{}
	for i := range respJson.Results {
		// Filter for movie or tv show results only.
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

	json.NewEncoder(w).Encode(transformResult)
}
