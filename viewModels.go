package main

type getMatchRequest struct {
	Id int64 `json:"id"`
}

type getMatchResponse struct {
	Match iMatch `json:"match"`
}

type createMatchRequest struct {
	GameId int64 `json:"gameId"`
}

type createMatchResponse struct {
	Id    int64  `json:"id"`
	Match iMatch `json:"match"`
}

type moveRequest struct {
	Id   int64 `json:"id"`
	Move iMove `json:"move"`
}

type moveResponse struct {
	Match iMatch `json:"match"`
}
