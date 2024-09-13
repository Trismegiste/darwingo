package main

import (
	"math/rand"
)

func explodingDice(side int) int {
	oneDie := rand.Intn(side) + 1
	roll := oneDie

	for oneDie == side {
		oneDie = rand.Intn(side) + 1
		roll += oneDie
	}

	return roll
}
