package services

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetMatch(request GetMatchRequest) GetMatchResponse {
	match, err := FindMatch(request.Id)
	if err != nil {
		panic(err)
	}
	return GetMatchResponse{
		Match: match,
	}
}

func CreateMatch(request CreateMatchRequest) CreateMatchResponse {
	boardSize := 8
	match := QuoridorMatch{
		Id:    primitive.NewObjectID(),
		Board: make([][]int, boardSize*2-1),
	}
	err := InsertMatch(match)
	if err != nil {
		panic(err)
	}
	return CreateMatchResponse{
		Match: match,
	}
}

func Move(request MoveRequest) MoveResponse {
	return MoveResponse{}
}
