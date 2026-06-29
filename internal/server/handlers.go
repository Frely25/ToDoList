package server

import (
	"ToDoList/internal/dto"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	// Проверка на метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Прочитать тело запроса и превратить его в объект нашей структуры
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error in request body", http.StatusBadRequest)
		return
	}

	req := dto.CreateTaskRequest{}

	if err := json.Unmarshal(httpRequestBody, &req); err != nil {
		http.Error(w, "Error in request body", http.StatusBadRequest)
		return
	}
	newTask, err := s.service.AddTask(req.Title, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	bytes, err := json.Marshal(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (s *Server) handleGetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	tasks := s.service.GetTasks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (s *Server) handleGetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	arrayOfPath := strings.Split(r.URL.Path, "/")
	num, err := strconv.Atoi(arrayOfPath[len(arrayOfPath)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task, err := s.service.GetTaskToID(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (s *Server) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	// Проверить метод
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Пропарсить строку и вытянуть id
	arrayOfPath := strings.Split(r.URL.Path, "/")
	num, err := strconv.Atoi(arrayOfPath[len(arrayOfPath)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// По id через service удалить таск
	deletedTask, err := s.service.DeleteTask(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Вернуть ответ пользоателю
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedTask)
}

func (s *Server) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	// Проверить метод
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Прочитать тело и за анмаршалить
	req := dto.UpdateTaskRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	arrayOfPath := strings.Split(r.URL.Path, "/")
	num, err := strconv.Atoi(arrayOfPath[len(arrayOfPath)-1])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.ID = num
	// Изменить наш таск
	newTask, err := s.service.UdpateTask(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Вернуть ответ с новым таском
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTask)
}
