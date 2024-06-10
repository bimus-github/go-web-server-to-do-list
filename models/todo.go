package models

type ToDo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
