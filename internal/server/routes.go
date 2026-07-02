package server

import (
	"errors"
	"net/http"
	"slices"
)

func (s *Server) RegisterRoutes() error {
	routes := map[string]func(w http.ResponseWriter, r *http.Request){
		"GET /tasks":         s.handleGetTasks,
		"GET /tasks/{id}":    s.handleGetTask,
		"POST /tasks/":       s.handleCreateTask,
		"DELETE /tasks/{id}": s.handleDeleteTask,
		"PATCH /tasks/{id}":  s.handleUpdateTask,
		"GET /histories":     s.handleGetHistory,
		"DELETE /histories":  s.handleClearHistory,
		"DELETE /tasks":      s.handleClearTasksOfList,
	}

	for pattern, handle := range routes {
		if err := s.AddHandle(pattern, handle); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) AddHandle(pattern string, function func(http.ResponseWriter, *http.Request)) error {
	if slices.Contains(s.patterns, pattern) {
		return errors.New("Такой паттерн уже есть")
	}
	http.HandleFunc(pattern, function)
	s.patterns = append(s.patterns, pattern)
	return nil
}
