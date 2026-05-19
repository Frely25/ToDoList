package task

type Storage interface {
	Save([]Task, int) error
	Load() ([]Task, int, error) // Задачи, счетсчик ID, ошибка
}
