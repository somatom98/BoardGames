package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/match/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getMatchRequest := getMatchRequest{
				Id: 1,
			}
			response, err := json.Marshal(getMatch(getMatchRequest))
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, string(response))
		} else if r.Method == "POST" {
			requestBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			var createMatchRequest createMatchRequest
			err = json.Unmarshal(requestBody, &createMatchRequest)
			response, err := json.Marshal(createMatch(createMatchRequest))
			fmt.Fprintf(w, string(response))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMatch(request getMatchRequest) getMatchResponse {
	boardSize := 8
	return getMatchResponse{
		Match: quoridorMatch{
			Id:    1,
			Board: make([][]int, boardSize*2-1),
		},
	}
}

func createMatch(request createMatchRequest) createMatchResponse {
	boardSize := 8
	return createMatchResponse{
		Id: 1,
		Match: quoridorMatch{
			Id:    1,
			Board: make([][]int, boardSize*2-1),
		},
	}
}

func move(request moveRequest) moveResponse {
	return moveResponse{}
}
