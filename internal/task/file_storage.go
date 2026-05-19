package task

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileStorage struct {
	path_to_tasks  string
	path_to_nextId string
}

const path_to_data string = "data/"

func NewFileStorage(path_to_tasks, path_to_nextId string) *FileStorage {
	// Проверить существуют ли такие файлы
	if path_to_tasks == "" {
		path_to_tasks = path_to_data + "tasks.json"
	}
	if path_to_nextId == "" {
		path_to_nextId = path_to_data + "config.txt"
	}
	// Я думаю, можно сделать через цикл:
	var file *os.File
	var err error
	for _, path := range [2]string{path_to_tasks, path_to_nextId} {
		file, err = os.Open(path)
		if err != nil {
			// Если нет, создать, если да - открыть
			file, err = os.Create(path)
			if err != nil {
				return nil
			}
		}
		defer file.Close()
	}
	return &FileStorage{ // Вернуть структуру с заполненым путем
		path_to_tasks:  path_to_tasks,
		path_to_nextId: path_to_nextId,
	}
}

func (file FileStorage) Save(task_to_save []Task, nextId_to_save int) error {
	// Marshal
	jsonObj, err := json.MarshalIndent(task_to_save, "", "	")
	if err != nil {
		return err
	}
	// запись в файл задач
	err = os.WriteFile(file.path_to_tasks, jsonObj, 0644)
	if err != nil {
		return err
	}
	//  запись в файл айдишника
	err = os.WriteFile(file.path_to_nextId, []byte(fmt.Sprintf("temp %d", nextId_to_save)), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (file FileStorage) Load() ([]Task, int, error) {
	var tasks_to_ret []Task
	// Открыть файл
	jsonObj, err := os.ReadFile(file.path_to_tasks)
	if err != nil {
		return []Task{}, 0, err
	}
	// Unmarshal
	err = json.Unmarshal(jsonObj, &tasks_to_ret)
	if err != nil {
		return []Task{}, 0, err
	}
	// считываем файл с nextId
	nextId_to_ret, err := file.loadNextId()
	if err != nil {
		return []Task{}, 0, nil
	}
	return tasks_to_ret, nextId_to_ret, nil // Возвращаем слайс
}

func (file FileStorage) loadNextId() (int, error) {
	obj, err := os.ReadFile(file.path_to_nextId)
	if err != nil {
		return -1, err
	}
	data_str := strings.TrimSpace(string(obj))
	return strconv.Atoi(data_str) // Возвращаем счетсчик для id
}
