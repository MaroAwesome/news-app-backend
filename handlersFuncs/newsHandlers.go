package handlersFuncs

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../controllers"
	"../models"
)

func CreateNewsEndpoint(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	var new models.News
	_ = json.NewDecoder(request.Body).Decode(&new)
	result := controllers.InsertNews(&new)
	json.NewEncoder(response).Encode(result)
}

func UpdateNewsEndpoint(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	var new models.News
	_ = json.NewDecoder(request.Body).Decode(&new)
	result := controllers.UpdateNew(&new)
	json.NewEncoder(response).Encode(result)
}

func GetAllNewsEndpoint(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")

	news, err := controllers.GetAllNews()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(news)
}

func GetNewsEndpoint(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	id := mux.Vars(request)["id"]
	news, err := controllers.GetNew(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(news)
}
