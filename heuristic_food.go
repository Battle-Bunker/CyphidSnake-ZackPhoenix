package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

/*
type GameSnapshot interface {
	GameID() string
	Rules() rules.Ruleset
	Turn() int
	Height() int
	Width() int
	Food() []rules.Point
	Hazards() []rules.Point
	Snakes() []SnakeSnapshot
	You() SnakeSnapshot
	Teammates() []SnakeSnapshot
	YourTeam() []SnakeSnapshot
	Opponents() []SnakeSnapshot
	ApplyMoves(moves []rules.SnakeMove) (GameSnapshot, error)
}

type SnakeSnapshot interface {
	ID() string
	Name() string
	Health() int
	Body() []rules.Point
	Head() rules.Point
	Length() int
	LastShout() string
	ForwardMoves() []rules.SnakeMove
}
*/

// TODO implement a heuristic that returns higher values when closer to food
func HeuristicFood(snapshot agent.GameSnapshot) int {
	return 0
}
