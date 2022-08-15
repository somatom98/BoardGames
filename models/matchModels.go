package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"n"`
}

type Board [][]int

type Coordinate struct {
	X int
	Y int
}

var UP = Coordinate{X: 0, Y: 1}
var DOWN = Coordinate{X: 0, Y: -1}
var RIGHT = Coordinate{X: 1, Y: 0}
var LEFT = Coordinate{X: -1, Y: 0}

type IMatch interface {
	GetId() primitive.ObjectID
	GetGameId() primitive.ObjectID
	GetBoard() Board
	NewBoard(playersNumber int) Board
	MakeMove(move IMove) (Board, error)
}

type IMove interface {
	IsValid(IMatch) error
}

type GetMatchRequest struct {
	Id string `json:"id"`
}

type GetMatchResponse struct {
	Match IMatch `json:"match"`
}

type CreateMatchRequest struct {
	GameId string `json:"gameId"`
}

type CreateMatchResponse struct {
	Match IMatch `json:"match"`
}

type GetGamesResponse struct {
	Games []Game `json:"games"`
}

type MoveRequest struct {
	MatchId string                 `json:"matchId"`
	Move    map[string]interface{} `json:"move"`
}

type MoveResponse struct {
	Match IMatch `json:"match"`
}
