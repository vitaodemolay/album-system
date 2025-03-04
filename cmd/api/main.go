package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitaodemolay/album-system/cmd/api/controller"
)

const (
	connectionString = "sqlserver://sa:PassW0rd@localhost:5433?database=SLQ_ALBUMSYSTEM_DB&connection+timeout=30"
)

func main() {
	log.Println("Starting server on port 8080...")
	router := mux.NewRouter()

	ctrl := controller.NewController(connectionString)
	router.HandleFunc("/", rootDefaultOK).Methods("GET")
	router.HandleFunc("/api/albums", ctrl.GetAlbums).Methods("GET")
	router.HandleFunc("/api/albums/{id}", ctrl.GetAlbumById).Methods("GET")
	router.HandleFunc("/api/albums", ctrl.CreateAlbum).Methods("POST")
	router.HandleFunc("/api/albums/{id}", ctrl.DeleteAlbum).Methods("DELETE")

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func rootDefaultOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Album API!"))
}
