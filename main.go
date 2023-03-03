package main

import (
	"net/http"

	"github.com/ahmadhamza/kg-backend/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.MainRouterPathHandeler).Methods("get")
	router.HandleFunc("/addTeacher", handlers.AddTeacher).Methods("post")

	http.ListenAndServe(":8080", router)
}
