package main

type game struct {
	Id   int64  `json:"id"`
	Name string `json:"title"`
}

type iMatch interface {
	makeMove(iMove) bool
}

type quoridorMatch struct {
	Id    int64   `json:"id"`
	Board [][]int `json:"board"`
}

func (match quoridorMatch) makeMove(move iMove) bool {
	return true
}

type iMove interface {
}
