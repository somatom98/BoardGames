package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/somatom98/board-games/services"
)

func GetMatch(w http.ResponseWriter, r *http.Request) {
	sl := strings.Split(r.RequestURI, "/")
	id := sl[len(sl)-1]
	getMatchRequest := services.GetMatchRequest{
		Id: id,
	}
	getMatchResponse, err := services.GetMatch(getMatchRequest)
	if err != nil {
		panic(err)
	}
	response, err := json.MarshalIndent(getMatchResponse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(response))
}

func PostMatch(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var createMatchRequest services.CreateMatchRequest
	err = json.Unmarshal(requestBody, &createMatchRequest)
	if err != nil {
		panic(err)
	}
	createMatchRespone, err := services.CreateMatch(createMatchRequest)
	if err != nil {
		panic(err)
	}
	response, err := json.MarshalIndent(createMatchRespone, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(response))
}
