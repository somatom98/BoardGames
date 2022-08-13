package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/somatom98/board-games/api"
)

func main() {

	httpHandler := api.NewHttpHandler()

	httpHandler.AddEndpoint("GET", "/match/", func(w http.ResponseWriter, r *http.Request) {
		getMatchRequest := api.GetMatchRequest{
			Id: 1,
		}
		response, err := json.Marshal(api.GetMatch(getMatchRequest))
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(response))
	})

	httpHandler.AddEndpoint("POST", "/match/", func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var createMatchRequest api.CreateMatchRequest
		err = json.Unmarshal(requestBody, &createMatchRequest)
		response, err := json.Marshal(api.CreateMatch(createMatchRequest))
		fmt.Fprintf(w, string(response))
	})

	httpHandler.ListenAndServe(":8080")
}
