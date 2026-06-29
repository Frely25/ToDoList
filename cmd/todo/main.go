package main

import (
	"ToDoList/internal/server"
	"context"
	"fmt"
	"strings"
	"time"
)

//"fmt"

func menuForServer(srv *server.Server) {
	for {
		fmt.Print("Введите \"Выкл\", чтобы остановить сервер: ")
		ans := ""
		fmt.Scan(&ans)
		if strings.ToLower(ans) == "выкл" {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := srv.ShutDown(ctx); err != nil {
				fmt.Println("Ошибка при выключении сервера: ", err)
			} else {
				fmt.Println("Вы успешно остановили сервер!")
			}

			return
		}
	}
}

func main() {
	mainServer, err := server.NewServer("9425")
	if err != nil {
		fmt.Println("Произошла ошибка при включении сервера: ", err)
		return
	}
	err = mainServer.RegisterRoutes()
	if err != nil {
		fmt.Println("При регистрации роутеров произошла ошибка: ", err)
		return
	}
	fmt.Println("Сервер запущен!")

	go menuForServer(mainServer)

	if err = mainServer.ServerUp(); err != nil {
		fmt.Println("Ошибка при работе сервера: ", err)
		return
	}
}
