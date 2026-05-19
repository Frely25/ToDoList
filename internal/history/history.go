package history

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

type History struct {
	path string
}

const path_to_data string = "build/"

func NewHistory(path string) *History {
	if path == "" {
		path = "history.log"
	}
	path = path_to_data + path
	file, err := os.Open(path) // Проверяем, существует ли файл
	if err != nil {
		// Если файл не существует, создаем его
		file, err = os.Create(path)
		if err != nil {
			return nil // Возвращаем nil, если не удалось создать файл
		}
	}
	defer file.Close()
	return &History{path: path}
}

func (h History) Log(message string) error {
	// Логгируем в формате [время дата] действие объект
	file, err := os.OpenFile(h.path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("Failed to open history log: " + err.Error()) // Возвращаем ошибку, если не удалось открыть файл
	}
	defer file.Close()
	message = fmt.Sprintf("[%s] %s", time.Now().Format("02.01.2006 15:04"), message) // Добавляем к сообщению текущую дату и время
	if _, err := file.WriteString(message + "\n"); err != nil {                      // Пишем сообщение в файл и проверяем на ошибки
		return errors.New("Writing to history log failed: " + err.Error()) // Возвращаем ошибку, если не удалось записать в файл
	}
	return nil
}

func (h History) GetHistory() ([]string, error) {
	file, err := os.Open(h.path)
	if err != nil {
		return nil, errors.New("Failed to open history log: " + err.Error()) // Возвращаем ошибку, если не удалось открыть файл
	}
	defer file.Close()
	var history []string
	scanner := bufio.NewScanner(file) // Сканируем файл построчно и добавляем каждую строку в срез history
	for scanner.Scan() {              // Сканируем файл построчно
		history = append(history, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("Failed to read history log: " + err.Error()) // Возвращаем ошибку, если не удалось прочитать файл
	}
	return history, nil
}

func (h History) ClearHistory() error {
	if err := os.WriteFile(h.path, []byte("History is reset"), 0644); err != nil {
		return err
	}
	return nil
}
