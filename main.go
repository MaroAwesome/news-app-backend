package main

import (
	"fmt"
	"net/http"

	"./handlersFuncs"
	"./models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	models.Init()

	router.HandleFunc("/getAll", handlersFuncs.GetAllNewsEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/get/{id}", handlersFuncs.GetNewsEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/insertNews", handlersFuncs.CreateNewsEndpoint).Methods("POST")
	router.HandleFunc("/updateNews", handlersFuncs.UpdateNewsEndpoint).Methods("PUT", "OPTIONS")
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "DELETE", "GET"})
	fmt.Println("Server now running on port 8000")

	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router))
	// to enable hls uncommit the following and insert your cert and pk
	/*	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", handlers.CORS(originsOk, headersOk, methodsOk)(router))
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	*/
}
