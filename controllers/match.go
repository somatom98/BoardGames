package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/somatom98/board-games/services"
)

func GetMatch(w http.ResponseWriter, r *http.Request) {
	sl := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(sl[len(sl)-1])
	if err != nil {
		panic(err)
	}
	getMatchRequest := services.GetMatchRequest{
		Id: int64(id),
	}
	response, err := json.Marshal(services.GetMatch(getMatchRequest))
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
	response, err := json.Marshal(services.CreateMatch(createMatchRequest))
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(response))
}
