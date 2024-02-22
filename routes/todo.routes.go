package routes

import (
	"encoding/json"
	"net/http"

	"github.com/daniir/go_api/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Conn struct {
	db gorm.DB
}

func New(db gorm.DB) *Conn {
	return &Conn{
		db: db,
	}
}

func (c Conn) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	c.db.Model(&models.Todo{}).Find(&todos)
	json.NewEncoder(w).Encode(&todos)
}

func (c Conn) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	getModel := c.db.First(&todo, params["id"])
	if getModel.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(getModel.Error.Error()))
		return
	}

	if todo.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Item not found"))
		return
	}

	json.NewEncoder(w).Encode(&todo)
}

func (c Conn) CreateNewTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	createTodo := c.db.Create(&newTodo)
	if createTodo.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createTodo.Error.Error()))
	}

	json.NewEncoder(w).Encode(&newTodo)
}

func (c Conn) UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	params := mux.Vars(r)
	getModel := c.db.First(&todo, params["id"])

	if getModel.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(getModel.Error.Error()))
		return
	}

	if todo.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Item not found"))
	}

	c.db.Model(&todo).Update("status", true)
	json.NewEncoder(w).Encode(&todo)
}

func (c Conn) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	getModel := c.db.First(&todo, params["id"])

	if getModel.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(getModel.Error.Error()))
		return
	}

	if todo.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Item not found"))
	}

	c.db.Delete(&todo)
}
