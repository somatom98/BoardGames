package api

func GetMatch(request GetMatchRequest) GetMatchResponse {
	boardSize := 8
	return GetMatchResponse{
		Match: QuoridorMatch{
			Id:    1,
			Board: make([][]int, boardSize*2-1),
		},
	}
}

func CreateMatch(request CreateMatchRequest) CreateMatchResponse {
	boardSize := 8
	return CreateMatchResponse{
		Id: 1,
		Match: QuoridorMatch{
			Id:    1,
			Board: make([][]int, boardSize*2-1),
		},
	}
}

func Move(request MoveRequest) MoveResponse {
	return MoveResponse{}
}
