package dto

import "time"

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"descriprion"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"descriprion"`
	Completed   bool   `json:"completed"`
}

type ResponseTask struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	DateCreate    time.Time `json:"datacreate"`
	DateCompleted time.Time `json:"datacomp"`
	Completed     bool      `json:"completed"`
}
