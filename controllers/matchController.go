package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/somatom98/board-games/models"
	"github.com/somatom98/board-games/services"
)

func GetMatch(w http.ResponseWriter, r *http.Request) {
	sl := strings.Split(r.RequestURI, "/")
	id := sl[len(sl)-1]
	getMatchRequest := models.GetMatchRequest{
		Id: id,
	}
	getMatchResponse, err := services.GetMatch(getMatchRequest)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	response, err := json.MarshalIndent(getMatchResponse, "", "\t")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, string(response))
}

func PostMatch(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	var createMatchRequest models.CreateMatchRequest
	err = json.Unmarshal(requestBody, &createMatchRequest)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	createMatchRespone, err := services.CreateMatch(createMatchRequest)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	response, err := json.MarshalIndent(createMatchRespone, "", "\t")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, string(response))
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	getGamesResponse, err := services.GetGames()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	response, err := json.MarshalIndent(getGamesResponse, "", "\t")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, string(response))
}

func PostMove(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	var moveRequest models.MoveRequest
	err = json.Unmarshal(requestBody, &moveRequest)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	moveResponse, err := services.Move(moveRequest)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	response, err := json.MarshalIndent(moveResponse, "", "\t")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, string(response))
}
