package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-list/models"
	"todo-list/storage"

	"github.com/gorilla/mux"
)

type ToDoHandler struct {
	Storage *storage.MemoryStorage
}

func (h *ToDoHandler) GetToDos(w http.ResponseWriter, r *http.Request) {
	todos := h.Storage.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func (h *ToDoHandler) CreateToDo(w http.ResponseWriter, r *http.Request) {
	var todo models.ToDo
	json.NewDecoder(r.Body).Decode(&todo)
	created := h.Storage.Create(todo)
	json.NewEncoder(w).Encode(created)
}

func (h *ToDoHandler) GetToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if todo, exists := h.Storage.Get(id); exists {
		json.NewEncoder(w).Encode(todo)
	} else {
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: "Not found"})
	}
}

func (h *ToDoHandler) UpdateToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var todo models.ToDo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = id
	// print todo
	fmt.Println(todo)
	if h.Storage.Update(id, todo) {
		json.NewEncoder(w).Encode(todo)
	} else {
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: "Not found"})
	}
}

func (h *ToDoHandler) DeleteToDoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if h.Storage.Delete(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		json.NewEncoder(w).Encode(models.ErrorMessage{Message: "Not found"})
	}
}
