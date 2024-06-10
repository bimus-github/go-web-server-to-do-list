package storage

import (
	"sync"
	"todo-list/models"
)

type MemoryStorage struct {
	mu    sync.Mutex
	todos map[int]models.ToDo
	id    int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[int]models.ToDo),
	}
}

func (s *MemoryStorage) Create(todo models.ToDo) models.ToDo {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.id++
	todo.ID = s.id
	s.todos[todo.ID] = todo
	return todo
}

func (s *MemoryStorage) GetAll() []models.ToDo {
	s.mu.Lock()
	defer s.mu.Unlock()
	todos := make([]models.ToDo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (s *MemoryStorage) Get(id int) (models.ToDo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo, exists := s.todos[id]
	return todo, exists
}

func (s *MemoryStorage) Update(id int, todo models.ToDo) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.todos[id]; exists {
		s.todos[id] = todo
		return true
	}
	return false
}

func (s *MemoryStorage) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.todos[id]; exists {
		delete(s.todos, id)
		return true
	}
	return false
}
