package db

type Show struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Watched bool   `json:"watched"`
}
