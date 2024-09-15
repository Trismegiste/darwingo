package darwin

import (
	"fmt"
	"main/random"
	"math/rand"
	"slices"
)

var pool []Fighter
var poolSize int

// initialises all fighters
func Initialise(size int) {
	poolSize = size
	for range poolSize {
		pool = append(pool, BuildFighter(random.RandomTrait(),
			rand.Intn(3),
			random.RandomTrait(),
			random.RandomTrait()))
	}
}

// Runs one step in the evolution
func RunEpoch(maxRound int) {
	for idx := range poolSize {
		pool[idx].resetEpoch()
	}

	for f1 := range poolSize {
		for f2 := 0; f2 < f1; f2++ {
			runFight(&pool[f1], &pool[f2], maxRound)
		}
	}
}

// Runs one fight between two fighters
func runFight(fighter1, fighter2 *Fighter, maxRound int) {
	// initialise fight
	fighter1.resetFight()
	fighter2.resetFight()
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

	// group the pool per cost
	perCost := make(map[int][]*Fighter)
	for _, npc := range pool {
		cost := npc.getCost()
		perCost[cost] = append(perCost[cost], &npc)
	}

	// selection per cost
	for cost, group := range perCost {
		best := group[0]
		weaker := group[len(group)-1]
		fmt.Println("At cost:", cost, "we have", len(group), "fighters",
			"and the best is", best,
			"and the weaker is", weaker)

		weaker.clone(best)
		weaker.mutate()
	}
}

// Prints some info
func Log(howmany int) {
	for k := 0; k < howmany; k++ {
		fmt.Println(pool[k])
	}
}
