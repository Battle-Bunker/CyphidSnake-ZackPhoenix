package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"fmt"
	"log"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicHealth(snapshot agent.GameSnapshot) int {
	totalHealth := 0
	snakeStats := ""
	for _, snake := range snapshot.YourTeam() {
		totalHealth += snake.Health()
		snakeStats += fmt.Sprintf("[ID: %s, Health: %d] ", snake.ID(), snake.Health())
	}
	log.Printf("Turn %d - Snakes IDs and Health: %s\n", snapshot.Turn(), snakeStats)
	return totalHealth
}
