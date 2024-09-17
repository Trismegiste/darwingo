package darwin

import (
	"fmt"
	"main/random"
	"math/rand"
	"slices"
)

type World struct {
	pool    []*Fighter
	perCost map[int][]*Fighter
}

// initialises all fighters
func BuildWorld(size int) *World {
	w := new(World)
	for range size {
		w.pool = append(w.pool, BuildFighter(random.RandomTrait(),
			rand.Intn(3),
			random.RandomTrait(),
			random.RandomTrait(),
			random.RandomTrait()))
	}

	return w
}

// Runs one step in the evolution
func (w *World) RunEpoch(maxRound int) {
	// Reset
	for _, f := range w.pool {
		f.resetEpoch()
	}

	// Battle
	for f1 := range len(w.pool) {
		for f2 := 0; f2 < f1; f2++ {
			runFight(w.pool[f1], w.pool[f2], maxRound)
		}
	}

	// Stat on current epoch
	slices.SortFunc(w.pool, func(a, b *Fighter) int {
		return b.victory - a.victory
	})

	// Groups the pool by cost of fighter
	w.perCost = make(map[int][]*Fighter)
	for _, npc := range w.pool {
		cost := npc.getCost()
		w.perCost[cost] = append(w.perCost[cost], npc)
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
func (w *World) Selection() {
	// selection per cost
	for cost, group := range w.perCost {
		best := group[0]
		weaker := group[len(group)-1]
		fmt.Println("At cost:", cost, "we have", len(group), "fighters",
			"and the best is", best,
			"and the weaker is", weaker)

		weaker.mimic(best)
		weaker.mutate()
	}
}

// Prints some info
func (w *World) Log(howmany int) {
	for k := 0; k < howmany; k++ {
		fmt.Println(w.pool[k])
	}
}

type StatCost struct {
	GroupCount  int
	AvgVictory  float32
	BestFighter *Fighter
}

func (w *World) GetStatPerCost() map[int]*StatCost {
	// the stats we'll return
	stat := make(map[int]*StatCost)

	for cost, group := range w.perCost {
		info := new(StatCost)
		info.BestFighter = group[0]
		info.GroupCount = len(group)
		sum := 0
		for _, fighter := range group {
			sum += fighter.victory
		}
		info.AvgVictory = float32(sum) / float32(len(group))
		stat[cost] = info
	}

	return stat
}
