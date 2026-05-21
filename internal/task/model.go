package task

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func (task Task) GetID() int {
	return task.ID
}

func (task Task) GetCompleted() bool {
	return task.Completed
}
