package ui

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readInt() (int, error) {
	input, err := reader.ReadString('\n') // Читаем строку до символа новой строки
	if err != nil {                       // Если произошла ошибка при чтении, возвращаем её
		return 0, err
	}
	input = strings.TrimSpace(input)                   // Удаляем пробелы и символы новой строки
	if value, err := strconv.Atoi(input); err == nil { // Преобразуем строку в целое число
		return value, nil // Если преобразование прошло успешно, возвращаем число и nil (нет ошибки)
	}
	return 0, errors.New("invalid input: not an integer") // Если преобразование не удалось, возвращаем 0 и ошибку
}

func readLine() string {
	input, err := reader.ReadString('\n') // Читаем строку до символа новой строки
	if err != nil {                       // Если произошла ошибка при чтении, возвращаем пустую строку
		return ""
	}
	return strings.TrimSpace(input) // Удаляем пробелы и символы новой строки и возвращаем результат
}
