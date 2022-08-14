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
	gameId, err := primitive.ObjectIDFromHex(request.GameId)
	if err != nil {
		return CreateMatchResponse{}, err
	}
	game, err := FindGame(gameId)
	if err != nil {
		return CreateMatchResponse{}, nil
	}
	match := newMatch(game)
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

func newMatch(game Game) IMatch {
	boardSize := 8
	match := QuoridorMatch{
		Id:     primitive.NewObjectID(),
		GameId: game.Id,
		Board:  make([][]int, boardSize*2-1),
	}
	for i := range match.Board {
		match.Board[i] = make([]int, boardSize*2-1)
	}
	return match
}
