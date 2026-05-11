package ui

import (
	"ToDoList/internal/task"
	"fmt"
)

var ser task.Service

func Welcome() {
	ser = task.NewService()
	fmt.Print("Добро пожаловать в программу-помощник ToDoList!\n\n\n")
}

func Menu() {
	isLive := true   // Отвечает за работу программы
	var title string // Переменная для ввода названий
	for isLive {
		fmt.Print("Выберете пункт меню:\n" +
			"\t1 - Добавить задачу\n" +
			"\t2 - Отметить задачу\n" +
			"\t3 - Удалить задачу\n" +
			"\t4 - Показать задачи\n" +
			"\t5 - Посмотреть историю\n" +
			"\t6 - Выйти\n" +
			"--->")
		choose, err := ReadInt()
		if err != nil {
			fmt.Printf("Вы допустили ошибку!\nОшибка: %s", err)
		} else {
			switch choose {
			case 1:
				// Добавить задачу
				fmt.Print("Введите название задачи: ")
				title = ReadLine()
				if err := ser.AddTask(title); err != nil {
					fmt.Printf("Ошибка: %s\n\n", err)
				} else {
					fmt.Print("Задача успешно добавлена\n\n")
				}
			case 2:
				// Отметить задачу
				fmt.Print("Введите ID задачи, которую хотите отметить: ")
				id, errInt := ReadInt()
				if errInt != nil {
					fmt.Printf("Ошибка: %s\n\n", errInt)
				} else {
					if err := ser.CompleteTask(id); err != nil {
						fmt.Printf("Ошибка: %s\n\n", err)
					} else {
						fmt.Print("Задача успешно отмечена\n\n")
					}
				}
			case 3:
				// Удалить задачу
				fmt.Print("Введите ID задачи, которую хотите удалить: ")
				id, errInt := ReadInt()
				if errInt != nil {
					fmt.Printf("Ошибка: %s\n\n", errInt)
				} else {
					if err := ser.DeleteTask(id); err != nil {
						fmt.Printf("Ошибка: %s\n\n", err)
					} else {
						fmt.Print("Задача успешно удалена\n\n")
					}
				}
			case 4:
				// Показать список задач
				fmt.Print("Список всех задач: \n")
				fmt.Print("Список всех задач: \n")
				fmt.Print("Список всех задач: \n")
			case 5:
				// Посмотреть изменения
			case 6:
				fmt.Println("Вы успешно вышли из программы")
				isLive = false
			default:
				fmt.Print("Вы ввели не существующий пункт меню!\n\n")
			}
		}
	}
}
