package models

type Todo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Descr string `json:"description"`
	Done  bool   `json:"done"`
}
