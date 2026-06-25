package task

import "time"

type Task struct {
	ID           int        // Айдишник
	Title        string     // Название
	Description  string     // Описание
	DateCreate   time.Time  // Дата создание
	DateComplete *time.Time // Дата Выполнения
	Completed    bool       // Выполнена ли
}

func (task Task) GetID() int {
	return task.ID
}

func (task Task) GetCompleted() bool {
	return task.Completed
}
