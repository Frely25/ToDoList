package task

type Task struct {
	ID    int    // Айдишник
	Title string // Название
	//  Description  // Описание
	//  DateCreate   // Дата создание
	//  DateComplete // Дата Выполнения
	Completed bool // Выполнена ли
}

func (task Task) GetID() int {
	return task.ID
}

func (task Task) GetCompleted() bool {
	return task.Completed
}
