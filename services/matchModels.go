package services

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"n"`
}

type IMatch interface {
	GetId() primitive.ObjectID
	GetGameId() primitive.ObjectID
	GetBoard() [][]int
	MakeMove(IMove) bool
}

type QuoridorMatch struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	GameId primitive.ObjectID `json:"game_id" bson:"g_id"`
	Board  [][]int            `json:"board" bson:"b"`
}

func (match QuoridorMatch) GetId() primitive.ObjectID {
	return match.Id
}

func (match QuoridorMatch) GetGameId() primitive.ObjectID {
	return match.Id
}

func (match QuoridorMatch) GetBoard() [][]int {
	return match.Board
}

func (match QuoridorMatch) MakeMove(move IMove) bool {
	return true
}

type IMove interface {
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

type MoveRequest struct {
	Id   string `json:"id"`
	Move IMove  `json:"move"`
}

type MoveResponse struct {
	Match IMatch `json:"match"`
}
