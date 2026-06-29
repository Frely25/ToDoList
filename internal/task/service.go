package task

import (
	"ToDoList/internal/dto"
	"ToDoList/internal/history"
	"errors"
	"fmt"
	"time"
)

type Service struct {
	tasks   []Task           // Задачи
	storage Storage          // Хранение данных (может быть что угодно)
	history *history.History // История изменений
	nextId  int              // Счетсчик для ID
}

func NewService() Service {
	storage := NewFileStorage("tasks.json", "config.txt")
	history := history.NewHistory("history.log")
	history.Log("Application is starting")
	tasks, nextId, error := storage.Load() // Загружаем все задачи из файла
	if error == nil {
		return Service{
			tasks:   tasks,
			storage: storage,
			history: history,
			nextId:  nextId,
		}
	} else {
		return Service{
			tasks:   []Task{},
			storage: storage,
			history: history,
			nextId:  0,
		}
	}
}

// Метод для добавления задачи
func (ser *Service) AddTask(title string, description string) (*Task, error) {
	// Проверяем не пустая ли новая задача
	if title != "" {
		// Добавляем задачу к слайсу
		newTask := Task{
			ID:           ser.nextId,
			Title:        title,
			Description:  description,
			DateCreate:   time.Now(),
			DateComplete: nil,
			Completed:    false,
		}
		ser.tasks = append(ser.tasks, newTask)

		mess := fmt.Sprintf("Add Task %d", ser.nextId) // Создаем сообщение для логгирования
		ser.history.Log(mess)                          // Логгируем
		ser.nextId++                                   // Добавляем к счетсчику 1

		return &newTask, nil
	} else {
		return nil, errors.New("Title is empty")
	}
}

// Метод для выполнения задачи
func (ser *Service) CompleteTask(id int) error {
	for i := 0; i < len(ser.tasks); i++ {
		// Проверям по id
		if ser.tasks[i].ID == id {
			ser.tasks[i].Completed = true // Помечаем задачу как выполненную
			now := time.Now()
			ser.tasks[i].DateComplete = &now
			mess := fmt.Sprintf("Complete Task %d", id) // Создаем сообщение для логгирования
			ser.history.Log(mess)                       // Логгируем
			return nil
		}
	}
	return errors.New("Index not found")
}

func (ser *Service) CompleteTaskToTask(task *Task) {
	task.Completed = true
	now := time.Now()
	task.DateComplete = &now
}

// Обновляем задачу по ID
func (ser *Service) UdpateTask(req dto.UpdateTaskRequest) (*Task, error) {
	for i := 0; i < len(ser.tasks); i++ {
		if ser.tasks[i].ID == req.ID {
			// Обновляем инфу
			temp := &ser.tasks[i]
			if req.Title == "" {
				return nil, errors.New("Title is empty")
			} else {
				temp.Title = req.Title
				temp.Description = req.Description

				if req.Completed {
					ser.CompleteTaskToTask(temp)
				} else {
					temp.Completed = req.Completed
				}
			}
			return temp, nil
		}
	}
	return nil, errors.New("Index not found")
}

// Метод для удаления задачи
func (ser *Service) DeleteTask(id int) (*Task, error) {
	for i := range ser.tasks {
		if ser.tasks[i].ID == id {
			ser.tasks = append(ser.tasks[:i], ser.tasks[i+1:]...) // Удаление элмента
			mess := fmt.Sprintf("Delete Task %d", id)             // Создаем сообщение для логгирования
			ser.history.Log(mess)                                 // Логгируем
			return &ser.tasks[i], nil
		}
	}
	return nil, errors.New("Index not found")
}

func (ser *Service) GetTasks() []Task {
	return ser.tasks
}

func (ser *Service) GetHistory() ([]string, error) {
	return ser.history.GetHistory()
}

func (ser *Service) ClearListOfTask() bool {
	if err := ser.storage.clearTasks(); err == nil {
		ser.tasks = []Task{}
		ser.nextId = 0
		ser.history.Log("List of tasks is clear")
		return true
	}
	return false
}

func (ser *Service) ClearHistory() error {
	if err := ser.history.ClearHistory(); err != nil {
		return err
	}
	return nil
}

func (ser *Service) Completion() error {
	if err := ser.storage.Save(ser.tasks, ser.nextId); err != nil {
		return err
	}
	ser.history.Log("Application has shut down")
	return nil
}

func (ser *Service) GetTaskToID(id int) (*Task, error) {
	for i := range ser.tasks {
		if ser.tasks[i].ID == id {
			return &ser.tasks[i], nil
		}
	}
	return &Task{}, errors.New("Task not found")
}
