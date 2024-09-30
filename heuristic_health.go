package main

import (
	"github.com/ZackBudai/cyphid-snake/agent"
)

// heuristicHealth calculates the sum of health for all snakes in your team,
// including the player's snake.
func HeuristicHealth(snapshot agent.GameSnapshot) float64 {
	totalHealth := 0
	for _, snake := range snapshot.YourTeam() {
		totalHealth += snake.Health()
	}
	return float64(totalHealth)
}
