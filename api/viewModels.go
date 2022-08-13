package api

type GetMatchRequest struct {
	Id int64 `json:"id"`
}

type GetMatchResponse struct {
	Match IMatch `json:"match"`
}

type CreateMatchRequest struct {
	GameId int64 `json:"gameId"`
}

type CreateMatchResponse struct {
	Id    int64  `json:"id"`
	Match IMatch `json:"match"`
}

type MoveRequest struct {
	Id   int64 `json:"id"`
	Move IMove `json:"move"`
}

type MoveResponse struct {
	Match IMatch `json:"match"`
}
