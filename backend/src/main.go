package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)


func main() {
	//define router
	r := mux.NewRouter()

	//api endpoints
	r.Handle("/todo", http.HandlerFunc(ToDoCreate)).Methods("POST")
	r.Handle("/todolist", http.HandlerFunc(TodoList)).Methods("GET")

	//define options
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	//start server
	log.Fatal(http.ListenAndServe(":8081", corsWrapper.Handler(r)))
}