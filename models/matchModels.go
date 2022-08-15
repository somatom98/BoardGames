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

type IMatch interface {
	GetId() primitive.ObjectID
	GetGameId() primitive.ObjectID
	GetBoard() Board
	MakeMove(IMove) (Board, error)
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
	MatchId string `json:"matchId"`
	Move    IMove  `json:"move"`
}

type MoveResponse struct {
	Match IMatch `json:"match"`
}
