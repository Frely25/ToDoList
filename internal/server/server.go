package server

import (
	"ToDoList/internal/task"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	//"slices"
	"strconv"
)

type Server struct {
	port     string
	patterns []string
	httpSrv  *http.Server
	service  *task.Service
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
		ser := task.NewService()

		return &Server{
			port:     port,
			patterns: make([]string, 1),
			httpSrv: &http.Server{
				ReadTimeout:  10 * time.Second,
				WriteTimeout: 10 * time.Second,
			},
			service: &ser,
		}, nil
	} else {
		// fmt.Println("Вы ввели не корректный порт")
		return nil, errors.New("Вы ввели не корректный порт")
	}
}

func (s *Server) ShutDown(ctx context.Context) error {
	if err := s.service.Completion(); err != nil {
		fmt.Println("\nПроизошла ошибка при сохранении: ", err)
		return err
	}
	return s.httpSrv.Shutdown(ctx)
}

func (s *Server) ServerUp() error {
	s.httpSrv.Addr = ":" + s.port
	return s.httpSrv.ListenAndServe()
}
