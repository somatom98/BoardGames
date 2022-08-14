package services

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetMatch(request GetMatchRequest) (GetMatchResponse, error) {
	match, err := FindMatch(request.Id)
	if err != nil {
		return GetMatchResponse{}, err
	}
	return GetMatchResponse{
		Match: match,
	}, nil
}

func CreateMatch(request CreateMatchRequest) (CreateMatchResponse, error) {
	boardSize := 8
	gameId, err := primitive.ObjectIDFromHex(request.GameId)
	if err != nil {
		return CreateMatchResponse{}, err
	}
	match := QuoridorMatch{
		Id:     primitive.NewObjectID(),
		GameId: gameId,
		Board:  make([][]int, boardSize*2-1),
	}
	err = InsertMatch(match)
	if err != nil {
		return CreateMatchResponse{}, err
	}
	return CreateMatchResponse{
		Match: match,
	}, nil
}

func Move(request MoveRequest) MoveResponse {
	return MoveResponse{}
}
