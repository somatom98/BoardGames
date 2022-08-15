package models

import (
	"errors"

	"github.com/somatom98/board-games/pkg/mmath"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuoridorMatch struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	GameId primitive.ObjectID `json:"game_id" bson:"g_id"`
	Board  Board              `json:"board" bson:"b"`
}

type QuoridorAction int

const (
	MOVE  QuoridorAction = 1
	FENCE QuoridorAction = 2
)

type QuoridorMove struct {
	Player int            `json:"player"`
	Action QuoridorAction `json:"action"`
	X      int            `json:"x"`
	Y      int            `json:"y"`
}

const (
	P1 = 1
	P2 = 2
	P3 = 3
	P4 = 4
	F  = 5
)

func (match QuoridorMatch) GetId() primitive.ObjectID {
	return match.Id
}

func (match QuoridorMatch) GetGameId() primitive.ObjectID {
	return match.Id
}

func (match QuoridorMatch) GetBoard() Board {
	return match.Board
}

func (match QuoridorMatch) MakeMove(mv IMove) (Board, error) {
	if err := mv.IsValid(match); err != nil {
		return nil, err
	} else {
		move := mv.(QuoridorMove)
		if move.Action == MOVE {
			match.Board.movePlayer(move.Player, move.X, move.Y)
		} else {
			match.Board.placeFence(move.X, move.Y)
		}
		return match.GetBoard(), nil
	}
}

func (move QuoridorMove) IsValid(match IMatch) error {
	if move.Action == MOVE {
		return match.GetBoard().isValidMovement(move.Player, move.X, move.Y)
	} else {
		return match.GetBoard().isValidPlacement(move.X, move.Y)
	}
	return nil
}

func (board Board) isValidMovement(player int, x int, y int) error {
	if x >= len(board) || x < 0 || y >= len(board) || y < 0 {
		return errors.New("move out of board")
	}

	xP, yP, err := board.getPlayerCurrentPosition(player)
	if err != nil {
		return err
	}

	if mmath.Diff(x, xP)+mmath.Diff(y, yP) != 1 {
		// validMovements := board.getValidMovements(x, y)
	}

	return nil
}

func (board Board) getValidMovements(x int, y int) []Coordinate {
	var validMovements []Coordinate
	// TODO
	return validMovements
}

func (board Board) isValidPlacement(x int, y int) error {
	if board.isOutOfBoard(x, y) {
		return errors.New("placement out of board")
	} else if x%2 != 0 || y%2 != 0 {
		return errors.New("fences cannot be placed here")
	} else if board[x][y] != 0 {
		return errors.New("there is no space here")
	}
	return nil
}

func (board Board) isOutOfBoard(x int, y int) bool {
	return x >= len(board) || x < 0 || y >= len(board) || y < 0
}

func (board Board) movePlayer(player int, x int, y int) error {
	xP, yP, err := board.getPlayerCurrentPosition(player)
	if err != nil {
		return err
	}
	board[xP][yP] = 0
	board[x][y] = player
	return nil
}

func (board Board) getPlayerCurrentPosition(player int) (int, int, error) {
	for i, row := range board {
		for j, square := range row {
			if square == player {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("player not found")
}

func (board Board) placeFence(x int, y int) error {
	board[x][y] = F
	return nil
}
