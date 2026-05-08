package history

type History struct {
	path string
}

func NewHistory(path string) *History {
	return &History{}
}

func (h History) Log(message string) {
	// Логгируем в формате [время дата] действие объект
}
