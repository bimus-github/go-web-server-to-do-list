package router

import (
	"net/http"
	"todo-list/handlers"
	"todo-list/storage"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	storage := storage.NewMemoryStorage()
	handler := &handlers.ToDoHandler{Storage: storage}

	r.HandleFunc("/", serveMainHTML)
	r.HandleFunc("/todos", handler.GetToDos).Methods("GET")
	r.HandleFunc("/todos", handler.CreateToDo).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.GetToDoByID).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.UpdateToDoByID).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", handler.DeleteToDoByID).Methods("DELETE")

	return r
}

func serveMainHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/main.html")
}
