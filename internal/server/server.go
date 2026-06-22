package server

import (
	"errors"
	"net/http"
	"slices"
	"strconv"
)

type Server struct {
	port     string
	patterns []string
}

func NewServer(port string) (*Server, error) {
	// 1) Число
	// 2) 0 - 65535 // 2 - байта
	port_int, err := strconv.Atoi(port)
	if err != nil {
		// fmt.Println("Вы передали некорректный порт: ", err)
		return nil, err
	}
	if 0 <= port_int && port_int <= 65535 {
		return &Server{
			port:     port,
			patterns: make([]string, 1),
		}, nil
	} else {
		// fmt.Println("Вы ввели не корректный порт")
		return nil, errors.New("Вы ввели не корректный порт")
	}
}

func (s *Server) ServerUp() error {
	// fmt.Println("Сервер запускается...")
	if err := http.ListenAndServe(":"+s.port, nil); err != nil {
		// fmt.Println("Не предвиденная ошибка на сервере: ", err)
		return err
	}
	return nil
}

func (s *Server) AddHandle(pattern string, function func(http.ResponseWriter, *http.Request)) error {
	if slices.Contains(s.patterns, pattern) {
		return errors.New("Такой паттерн уже есть")
	}
	http.HandleFunc(pattern, function)
	return nil
}
