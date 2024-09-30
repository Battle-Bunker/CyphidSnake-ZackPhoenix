package main

import (
	"github.com/ZackBudai/cyphid-snake/agent"
	"github.com/ZackBudai/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "",
		Color:      "#888888",
		Head:       "default",
		Tail:       "default",
	}

	portfolio := agent.NewPortfolio(
		agent.NewHeuristic(1.0, "team-health", HeuristicHealth),
		agent.NewHeuristic(1.0, "food", HeuristicFood),
	)

	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}
