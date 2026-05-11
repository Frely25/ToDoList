package task

import (
	"ToDoList/internal/history"
	"errors"
	"fmt"
)

type Service struct {
	tasks   []Task           // Задачи
	storage Storage          // Хранение данных (может быть что угодно)
	history *history.History // История изменений
	nextId  int              // Счетсчик для ID
}

func NewService() Service {
	storage := NewFileStorage("../../data/tasks.json")
	history := history.NewHistory("../../data/history.log")
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
func (ser *Service) AddTask(title string) error {
	// Проверяем не пустая ли новая задача
	if title != "" {
		// Добавляем задачу к слайсу
		ser.tasks = append(ser.tasks, Task{
			ID:        ser.nextId,
			Title:     title,
			Completed: false,
		})
		mess := fmt.Sprintf("Add Task %d", ser.nextId) // Создаем сообщение для логгирования
		ser.history.Log(mess)                          // Логгируем
		ser.nextId++                                   // Добавляем к счетсчику 1
		return nil
	} else {
		return errors.New("Title is empty")
	}
}

// Метод для выполнения задачи
func (ser *Service) CompleteTask(id int) error {
	for i := 0; i < len(ser.tasks); i++ {
		// Проверям по id
		if ser.tasks[i].ID == id {
			ser.tasks[i].Completed = true                       // Помечаем задачу как выполненную
			mess := fmt.Sprintf("Complete Task %d", ser.nextId) // Создаем сообщение для логгирования
			ser.history.Log(mess)                               // Логгируем
			return nil
		}
	}
	return errors.New("Index not found")
}

// Метод для удаления задачи
func (ser *Service) DeleteTask(id int) error {
	for i := range ser.tasks {
		if ser.tasks[i].ID == id {
			ser.tasks = append(ser.tasks[:i], ser.tasks[i+1:]...) // Удаление элмента
			mess := fmt.Sprintf("Delete Task %d", ser.nextId)     // Создаем сообщение для логгирования
			ser.history.Log(mess)                                 // Логгируем
			return nil
		}
	}
	return errors.New("Index not found")
}

func (ser *Service) GetTasks() []Task {
	return ser.tasks
}
