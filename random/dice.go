package random

import (
	"math/rand"
)

func ExplodingDice(side int) int {
	oneDie := rand.Intn(side) + 1
	roll := oneDie

	for oneDie == side {
		oneDie = rand.Intn(side) + 1
		roll += oneDie
	}

	return roll
}

func JokerRoll(side int) int {
	return max(ExplodingDice(side), ExplodingDice(6))
}
