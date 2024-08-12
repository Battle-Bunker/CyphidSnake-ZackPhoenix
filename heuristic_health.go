package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
// Calculates all of the health of all the agents in your team and returns it as an integer. (written by jacob)
func HeuristicHealth(snapshot agent.GameSnapshot) int {
	totalHealth := 0
	for _, snake := range snapshot.YourTeam() {
		totalHealth += snake.Health()
	}
	return totalHealth
}
