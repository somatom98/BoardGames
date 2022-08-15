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

func (match QuoridorMatch) MakeMove(mv IMove) error {
	if err := mv.IsValid(match); err != nil {
		return err
	} else {
		move := mv.(QuoridorMove)
		if move.Action == MOVE {
			match.Board.movePlayer(move.Player, move.X, move.Y)
		} else {
			match.Board.placeFence(move.X, move.Y)
		}
		return nil
	}
}

func (move QuoridorMove) IsValid(match IMatch) error {
	if move.Action == MOVE {
		return match.GetBoard().isValidMovement(move.Player, move.X, move.Y)
	} else {
		return match.GetBoard().isValidPlacement(move.X, move.Y)
	}
}

func (board Board) isValidMovement(player int, x int, y int) error {
	if board.isOutOfBoard(x, y) {
		return errors.New("move out of board")
	}

	if x%2 != 1 || y%2 != 1 {
		return errors.New("players cannot be placed here")
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
	player := Coordinate{X: x, Y: y}

	validMovements = append(validMovements, board.getValidMovementsInDirection(player, UP)...)
	validMovements = append(validMovements, board.getValidMovementsInDirection(player, DOWN)...)
	validMovements = append(validMovements, board.getValidMovementsInDirection(player, LEFT)...)
	validMovements = append(validMovements, board.getValidMovementsInDirection(player, RIGHT)...)

	return validMovements
}

func (board Board) getValidMovementsInDirection(pl Coordinate, dir Coordinate) []Coordinate {
	if board.isOutOfBoard(pl.X+dir.X*2, pl.Y+dir.Y*2) {
		return nil
	}
	if board[pl.X+dir.X][pl.Y+dir.Y] != F { // fence
		return nil
	}

	dest := Coordinate{X: pl.X + dir.X*2, Y: pl.Y + dir.Y*2}
	if board[dest.X][dest.Y] == 0 { // no player ahead
		return []Coordinate{dest}
	} // player ahead
	if board[dest.X+dir.X][dest.Y+dir.Y] == 0 && board[dest.X+dir.X*2][dest.Y+dir.Y*2] == 0 { // no fence or other player ahead of the player found
		return []Coordinate{Coordinate{X: dest.X + dir.X*2, Y: dest.Y + dir.Y*2}}
	}
	if board[dest.X+dir.X][dest.Y+dir.Y] == F { // fence ahead
		var validMovements []Coordinate
		var sideDir Coordinate
		if dir.X == 1 {
			sideDir = Coordinate{X: 0, Y: 1}
		} else {
			sideDir = Coordinate{X: 1, Y: 0}
		}
		if board.isOutOfBoard(dest.X+sideDir.X*2, dest.Y+sideDir.Y*2) && // inside board
			board[dest.X+sideDir.X][dest.Y+sideDir.Y] != F && // no fence and
			board[dest.X+sideDir.X*2][dest.Y+sideDir.Y*2] == 0 { // no player
			validMovements = append(validMovements, Coordinate{X: dest.X + sideDir.X*2, Y: dest.Y + sideDir.Y*2})
		}
		if board.isOutOfBoard(dest.X+sideDir.X*2, dest.Y+sideDir.Y*2) && // inside board
			board[dest.X-sideDir.X][dest.Y-sideDir.Y] != F && // no fence and
			board[dest.X-sideDir.X*2][dest.Y-sideDir.Y*2] == 0 { // no player
			validMovements = append(validMovements, Coordinate{X: dest.X - sideDir.X*2, Y: dest.Y - sideDir.Y*2})
		}
		return validMovements
	}
	return nil
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
