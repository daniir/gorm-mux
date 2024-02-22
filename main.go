package main

import (
	"fmt"
	"net/http"

	"github.com/daniir/go_api/database"
	"github.com/daniir/go_api/models"
	"github.com/daniir/go_api/routes"
	"github.com/daniir/go_api/settings"
	"github.com/gorilla/mux"
)

func main() {
	s, err := settings.New()
	if err != nil {
		panic("Error")
	}

	db, _ := database.DbConnection(s)
	todo := routes.New(*db)

	db.AutoMigrate(&models.Todo{})

	r := mux.NewRouter()
	r.HandleFunc("/", todo.GetTodos).Methods("GET")
	r.HandleFunc("/{id}", todo.GetTodo).Methods("GET")
	r.HandleFunc("/", todo.CreateNewTodo).Methods("POST")
	r.HandleFunc("/{id}", todo.UpdateTodoStatus).Methods("PUT")
	r.HandleFunc("/{id}", todo.DeleteTodo).Methods("DELETE")

	port := fmt.Sprintf(":%d", s.Port)

	http.ListenAndServe(port, r)
	fmt.Printf("Server at localhost:%d\n", s.Port)
}
