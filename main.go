package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/manakuro/golang-restfulapi-sample/models"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", getUsers)
	})

	fmt.Println("Server listen at http://localhost:8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	u := []models.User{
		{ID: 1, Name: "Taro", Age: 20},
		{ID: 2, Name: "Jiro", Age: 25},
	}
	respond(w, http.StatusOK, u)
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
