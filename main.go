package main

import (
	"PBPUTS/Controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rooms", Controller.GetAllRooms).Methods("GET")
	router.HandleFunc("/participants", Controller.RoomDetail).Methods("GET")

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
