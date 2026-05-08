package task

type Storage interface {
	Save([]Task) error
	Load() ([]Task, int, error) // Задачи, счетсчик ID, ошибка
}
