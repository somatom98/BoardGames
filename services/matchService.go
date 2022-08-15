package services

import (
	m "github.com/somatom98/board-games/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMatch(request m.GetMatchRequest) (m.GetMatchResponse, error) {
	match, err := FindMatch(request.Id)
	if err != nil {
		return m.GetMatchResponse{}, err
	}
	return m.GetMatchResponse{
		Match: match,
	}, nil
}

func CreateMatch(request m.CreateMatchRequest) (m.CreateMatchResponse, error) {
	gameId, err := primitive.ObjectIDFromHex(request.GameId)
	if err != nil {
		return m.CreateMatchResponse{}, err
	}
	game, err := FindGame(gameId)
	if err != nil {
		return m.CreateMatchResponse{}, nil
	}
	match := newMatch(game)
	err = InsertMatch(match)
	if err != nil {
		return m.CreateMatchResponse{}, err
	}
	return m.CreateMatchResponse{
		Match: match,
	}, nil
}

func GetGames() (m.GetGamesResponse, error) {
	games, err := FindGames()
	if err != nil {
		return m.GetGamesResponse{}, err
	}
	return m.GetGamesResponse{
		Games: games,
	}, nil
}

func Move(request m.MoveRequest) m.MoveResponse {
	return m.MoveResponse{}
}

func newMatch(game m.Game) m.IMatch {
	boardSize := 8
	match := m.QuoridorMatch{
		Id:     primitive.NewObjectID(),
		GameId: game.Id,
		Board:  make([][]int, boardSize*2-1),
	}
	for i := range match.Board {
		match.Board[i] = make([]int, boardSize*2-1)
	}
	return match
}
