package task

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	// Проверить существует ли такой файл
	// Если нет, создать, если да - открыть
	// Вернуть структуру с заполненым путем
	return &FileStorage{}
}

func (file FileStorage) Save(task_to_save []Task) error {
	// Marshal
	// запись в файл
	return nil
}

func (file FileStorage) Load() ([]Task, int, error) {
	// Открыть файл
	// Unmarshal
	// Запсь в слайс
	return []Task{}, 0, nil // Возвращаем слайс
}

func (file FileStorage) LoadNextId() (int, error) {
	return 0, nil // Возвращаем счетсчик для id
}
