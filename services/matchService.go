package services

import (
	"errors"

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

func Move(request m.MoveRequest) (m.MoveResponse, error) {
	match, err := FindMatch(request.MatchId)
	if err != nil {
		return m.MoveResponse{}, err
	}
	game, err := FindGame(match.GetGameId())
	if err != nil {
		return m.MoveResponse{}, err
	}
	move, err := castMove(request.Move, game)
	if err != nil {
		return m.MoveResponse{}, err
	}
	board, err := match.MakeMove(move)
	if err != nil {
		return m.MoveResponse{}, err
	}
	if err := UpdateMatch(match.GetId(), board); err != nil {
		return m.MoveResponse{}, err
	}
	return m.MoveResponse{
		Match: match,
	}, err
}

func newMatch(game m.Game) m.IMatch {
	var match m.IMatch
	switch game.Name {
	case "Quoridor":
		quoridorMatch := m.QuoridorMatch{
			Id:     primitive.NewObjectID(),
			GameId: game.Id,
		}
		quoridorMatch.Board = quoridorMatch.NewBoard(2)
		match = quoridorMatch
	}
	return match
}

func castMove(moveToCast map[string]interface{}, game m.Game) (m.IMove, error) {
	switch game.Name {
	case "Quoridor":
		move := m.QuoridorMove{
			Player: int(moveToCast["player"].(float64)),
			Action: m.QuoridorAction(moveToCast["action"].(float64)),
			X:      int(moveToCast["x"].(float64)),
			Y:      int(moveToCast["y"].(float64)),
		}
		return move, nil
	}
	return nil, errors.New("game not found")
}
