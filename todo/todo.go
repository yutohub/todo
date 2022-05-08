package todo

type Todo struct {
	ID   int64  `json:"id"`
	Task string `json:"task"`
	Done int64  `json:"done"`
}
