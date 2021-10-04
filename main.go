package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":1010", router))
}

func handleHome(writer http.ResponseWriter, request *http.Request)  {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	var status int
	status = 200
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":status,
		"message":"Hello World",
	})
}
