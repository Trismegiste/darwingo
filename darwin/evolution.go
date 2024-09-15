package darwin

import (
	"fmt"
	"math/rand"
	"slices"
)

var pool []Fighter
var poolSize int

// initialises all fighters
func Initialise(size int) {
	poolSize = size
	for k := 0; k < poolSize; k++ {
		pool = append(pool, BuildFighter(4+2*rand.Intn(5), rand.Intn(3), 4+2*rand.Intn(5)))
	}
}

// Rubs one step in the evolution
func RunEpoch(maxRound int) {
	for f1 := 0; f1 < poolSize; f1++ {
		for f2 := 0; f2 < f1; f2++ {
			runFight(&pool[f1], &pool[f2], maxRound)
		}
	}
}

// Runs one fight between two fighters
func runFight(fighter1, fighter2 *Fighter, maxRound int) {
	// initialise fight
	fighter1.reset()
	fighter2.reset()
	round := 0

	// fight
	for !fighter1.isDead() && !fighter2.isDead() && round < maxRound {
		fighter1.receiveAttack(fighter2)
		fighter2.receiveAttack(fighter1)
		round++
	}

	// aftermath
	if !fighter1.isDead() && fighter2.isDead() {
		fighter1.incVictory()
	}
	if !fighter2.isDead() && fighter1.isDead() {
		fighter2.incVictory()
	}
}

// Darwin selection
func Selection() {
	// stat on current epoch
	slices.SortFunc(pool, func(a, b Fighter) int {
		return b.victory - a.victory
	})
}

// Prints some info
func Log(howmany int) {
	for k := 0; k < howmany; k++ {
		fmt.Println(pool[k])
	}
}
