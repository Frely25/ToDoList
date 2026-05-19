package task

import (
	"encoding/json"
	"os"
)

type FileStorage struct {
	path string
}

const path_to_data string = "../../data/"

func NewFileStorage(path string) *FileStorage {
	// Проверить существует ли такой файл
	if path == "" {
		path = path_to_data + "tasks.json"
	}
	file, err := os.Open(path_to_data + path)
	if err != nil {
		// Если нет, создать, если да - открыть
		file, err = os.Create(path_to_data + path)
		if err != nil {
			return nil
		}

	}
	defer file.Close()
	return &FileStorage{ // Вернуть структуру с заполненым путем
		path: path_to_data + path,
	}
}

func (file FileStorage) Save(task_to_save []Task) error {
	// Marshal
	jsonObj, err := json.MarshalIndent(task_to_save, "", "	")
	if err != nil {
		return err
	}
	// запись в файл
	err = os.WriteFile(path_to_data+file.path, jsonObj, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (file FileStorage) Load() ([]Task, int, error) {
	// Открыть файл
	// Unmarshal
	// Запсь в слайс
	return []Task{}, 0, nil // Возвращаем слайс
}

func (file FileStorage) LoadNextId() (int, error) {
	// file, err := os.OpenFile(path_to_data+"config.txt", os.O_RDWR, 0644)
	// if err != nil {
	// 	return 0, err
	// }
	return 0, nil // Возвращаем счетсчик для id
}
