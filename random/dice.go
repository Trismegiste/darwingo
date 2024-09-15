package random

import (
	"math/rand"
)

// Rolls one die that can explode (the max number of the die means it is re-rolled and added)
func ExplodingDice(side int) int {
	oneDie := rand.Intn(side) + 1
	roll := oneDie

	for oneDie == side {
		oneDie = rand.Intn(side) + 1
		roll += oneDie
	}

	return roll
}

// A Joker rolls her die alongside with an exploding d6. She keeps the max of the two dice
func JokerRoll(side int) int {
	return max(ExplodingDice(side), ExplodingDice(6))
}

func RandomTrait() int {
	return 4 + 2*rand.Intn(5)
}
