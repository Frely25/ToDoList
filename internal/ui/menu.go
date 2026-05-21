package ui

import (
	"ToDoList/internal/task"
	"fmt"
	"strings"
)

var ser task.Service

func Init_to_service() {
	welcome()
	ser = task.NewService()
	menu()
}

func welcome() {
	fmt.Print("Добро пожаловать в программу-помощник ToDoList!\n\n\n")
}

func menu() {
	isLive := true   // Отвечает за работу программы
	var title string // Переменная для ввода названий
	for isLive {
		fmt.Print("Выберете пункт меню:\n" +
			"\t1 - Добавить задачу\n" +
			"\t2 - Отметить задачу\n" +
			"\t3 - Удалить задачу\n" +
			"\t4 - Показать задачи\n" +
			"\t5 - Посмотреть историю\n" +
			"\t6 - Очистить полностью список задач\n" +
			"\t7 - Очистить логи\n" +
			"\t8 - Выйти\n" +
			"--->")
		choose, err := readInt()
		if err != nil {
			fmt.Printf("Вы допустили ошибку!\nОшибка: %s", err)
		} else {
			switch choose {
			case 1:
				// Добавить задачу
				fmt.Print("Введите название задачи: ")
				title = readLine()
				if err := ser.AddTask(title); err != nil {
					fmt.Printf("Ошибка: %s\n\n", err)
				} else {
					fmt.Print("Задача успешно добавлена\n\n")
				}
			case 2:
				// Отметить задачу
				if len(ser.GetTasks()) == 0 {
					fmt.Print("Список задач - \033[1m пуст \033[0m")
				} else {
					fmt.Print("Введите ID задачи, которую хотите отметить: ")
					id, errInt := readInt()
					if errInt != nil {
						fmt.Printf("Ошибка: %s\n\n", errInt)
					} else {
						task, err := ser.GetTaskToID(id)
						if err != nil {
							fmt.Printf("Error: %s", err)
						} else {
							if task.GetCompleted() {
								fmt.Println("Задача уже отмечена")
							} else {
								ser.CompleteTask(id)
								fmt.Print("Задача успешно отмечена\n\n")
							}
						}
					}
				}
			case 3:
				// Удалить задачу
				if len(ser.GetTasks()) == 0 {
					fmt.Print("Список задач - \033[1m пуст \033[0m")
				} else {
					fmt.Print("Введите ID задачи, которую хотите удалить: ")
					id, errInt := readInt()
					if errInt != nil {
						fmt.Printf("Ошибка: %s\n\n", errInt)
					} else {
						if err := ser.DeleteTask(id); err != nil {
							fmt.Printf("Ошибка: %s\n\n", err)
						} else {
							fmt.Print("Задача успешно удалена\n\n")
						}
					}
				}
			case 4:
				// Показать список задач
				if len(ser.GetTasks()) == 0 {
					fmt.Print("Список задач - \033[1m пуст \033[0m")
				} else {
					fmt.Print("Список всех задач: \n" +
						"ID\tTitle\t\t\tCompleted\n")
					for _, value := range ser.GetTasks() {
						fmt.Println("==========================================")
						fmt.Printf("%d\t", value.ID)
						fmt.Printf("%s\t\t\t", value.Title)
						sm := "✗"
						if value.Completed {
							sm = "✓"
						}
						fmt.Println(sm)
					}
				}
				fmt.Print("\n\n")
			case 5:
				// Посмотреть изменения
				fmt.Print("История изменений: \n")
				history, err := ser.GetHistory()
				if err != nil {
					fmt.Printf("Ошибка: %s\n\n", err)
				} else if len(history) == 0 {
					fmt.Print("Список логов - " + "\033[1m" + "пуст" + "\033[0m")
				} else {
					for _, value := range history {
						fmt.Println("==========================================")
						fmt.Println(value)
					}
				}
				fmt.Print("\n\n")
			case 6:
				//Очистить список задач
				if len(ser.GetTasks()) == 0 {
					fmt.Print("Список задач - \033[1m пуст \033[0m")
				} else {
					fmt.Print("Вы хотите удалить все сохраненные задачи?\n" +
						"Удалить все задачи без возможности восстановления (Y/N): ")
					answer := strings.ToLower(readLine())
					if answer == "y" {
						isAgree := false
						for i := 0; i < 5; i++ {
							fmt.Print("Для подтвержедения удаления введите \"Удалить все\" или \"Отмена\" для отмены удаления\n--->")
							answer = readLine()
							if strings.ToLower(answer) == "отмена" {
								break
							} else if answer == "Удалить все" {
								isAgree = true
								break
							}
						}

						if isAgree {
							if ser.ClearListOfTask() {
								fmt.Print("Вы успешно полностью очистили список задач")
							} else {
								fmt.Print("Что-то пошло не так, если все же хотите удалить список задач, то попробуйте перезайти в приложение и заного попробовать")
							}
						}
						fmt.Print("\n\n")
					}
				}
			case 7:
				// Очистить историю
				fmt.Print("Вы хотите очистить историю изменений?(Y/N)\n" +
					"--->")
				ans := strings.ToLower(readLine())
				if ans == "y" {
					if ser.ClearHistory() {
						fmt.Print("История успешно очищена")
					} else {
						fmt.Print("Что-то пошло не так, ререзапустить программу и еще раз попробуйте!")
					}
				}
				fmt.Print("\n\n")
			case 8:
				fmt.Println("Вы успешно вышли из программы")
				isLive = false

			default:
				fmt.Print("Вы ввели не существующий пункт меню!\n\n")
			}
		}
	}
	ser.Completion()
}
