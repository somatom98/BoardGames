package main

import (
	"github.com/somatom98/board-games/controllers"
	"github.com/somatom98/board-games/pkg/api"
)

func main() {

	httpHandler := api.NewHttpHandler(8080)

	httpHandler.AddEndpoint("GET", "/match/", controllers.GetMatch)
	httpHandler.AddEndpoint("POST", "/match", controllers.PostMatch)
	httpHandler.AddEndpoint("GET", "/games/", controllers.GetGames)

	httpHandler.ListenAndServe()
}
