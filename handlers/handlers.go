package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ahmadhamza/kg-backend/database"
	"github.com/ahmadhamza/kg-backend/models"
)

func MainRouterPathHandeler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "hello World")
}

func AddTeacher(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
	}

	// Decode the JSON object
	var teacher models.Teacher
	err = json.Unmarshal(body, &teacher)

	if err != nil {
		http.Error(w, "Error decode request body", http.StatusBadRequest)
	}

	collection, err := database.ConnectToDB()

	if err != nil {
		http.Error(w, "Could not connect to the database", http.StatusInternalServerError)
	}

	collection.InsertOne(context.Background(), teacher)

	res, err := json.Marshal(teacher)

	fmt.Printf("%+v\n", teacher)
	w.Write([]byte(res))
}
