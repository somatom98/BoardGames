package api

type Game struct {
	Id   int64  `json:"id"`
	Name string `json:"title"`
}

type IMatch interface {
	MakeMove(IMove) bool
}

type QuoridorMatch struct {
	Id    int64   `json:"id"`
	Board [][]int `json:"board"`
}

func (match QuoridorMatch) MakeMove(move IMove) bool {
	return true
}

type IMove interface {
}
