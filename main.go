package main

import (
	"fmt"
	"net/http"

	"./handlers"
	"./models"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	models.Init()

	router.HandleFunc("/getAll", handlers.GetAllNewsEndpoint).Methods("GET")
	router.HandleFunc("/get/{id}", handlers.GetNewsEndpoint).Methods("GET")
	router.HandleFunc("/insertNews", handlers.CreateNewsEndpoint).Methods("POST")
	router.HandleFunc("/updateNews", handlers.UpdateNewsEndpoint).Methods("PUT")

	fmt.Println("Server now running on port 8000")

	http.ListenAndServe(":8000", router)
	// to enable hls uncommit the following and insert your cert and pk 
/*	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	*/
}
