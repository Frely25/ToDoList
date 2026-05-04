package history

import (
	"ToDoList/internal/task"
	"fmt"
)

// Константы для действий в истории
const (
	ActionCreate = "CREATE"
	ActionUpdate = "UPDATE"
	ActionDelete = "DELETE"
)

// Log представляет собой запись в истории изменений задач
type Log struct {
	Time   string
	Action string
	Task   task.Task
}

// History хранит все записи изменений задач за запуск приложения
type History struct {
	Logs []Log
}

// Методы для работы с историей изменений задач
// AddLog добавляет новую запись в историю изменений задач
func (h *History) AddLog(action string, task task.Task) {
	log := Log{
		Time:   "2024-06-01T12:00:00Z", // Измните на реальное время !!!
		Action: action,
		Task:   task,
	}
	h.Logs = append(h.Logs, log)
}

// GetLogs возвращает все записи истории в виде строк
func (h History) getLogsToString() []string {
	logs := make([]string, len(h.Logs))
	for i, log := range h.Logs {
		logs[i] = fmt.Sprintf("[%s] %s: Task ID %d - %s", log.Time, log.Action, log.Task.ID, log.Task.Title)
	}
	return logs
}
