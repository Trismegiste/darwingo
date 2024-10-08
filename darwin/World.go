package darwin

import (
	"encoding/json"
	"fmt"
	"main/random"
	"math/rand"
	"os"
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
			random.RandomTrait(),
			rand.Intn(4),
			rand.Intn(2),
			random.RandomTrait(),
			rand.Intn(3),
			rand.Intn(2),
			rand.Intn(3),
			rand.Intn(3),
			rand.Intn(2),
		))
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

func runRound(fighter1, fighter2 *Fighter) {
	init1 := fighter1.getInitiative()
	init2 := fighter2.getInitiative()

	// default behavior, fighter1 attacks first, if fighter2 has initiative, we swap fighter1 and 2
	var a, b *Fighter
	if init2 > init1 {
		// fighter 2 has init, fighter1 is the first to receive attack from fighter2
		a = fighter1
		b = fighter2
	} else {
		b = fighter1
		a = fighter2
	}

	a.receiveAttack(b) // b is attacking a
	b.receiveAttack(a) // a is attacking b
}

// Runs one fight between two fighters
func runFight(fighter1, fighter2 *Fighter, maxRound int) {
	// initialise fight
	fighter1.resetFight()
	fighter2.resetFight()
	round := 0

	// fight
	for !fighter1.isDead() && !fighter2.isDead() && round < maxRound {
		fighter1.resetRound()
		fighter2.resetRound()
		runRound(fighter1, fighter2)
		round++
	}

	// aftermath: is someone is dead ?
	if !fighter1.isDead() && fighter2.isDead() {
		fighter1.incVictory()
	}
	if !fighter2.isDead() && fighter1.isDead() {
		fighter2.incVictory()
	}
	// aftermath : if maxRound is reached, who is the most wounded ?
	if round == maxRound {
		if fighter1.wounds > fighter2.wounds {
			fighter2.incVictory()
		}
		if fighter1.wounds < fighter2.wounds {
			fighter1.incVictory()
		}
	}
}

// Darwin selection : for each group by cost, we replace the loser by the best with mutation
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

type GroupInfo struct {
	GroupCount int
	Winner     *Fighter
	Loser      *Fighter
}

type CostStat struct {
	InfoPerCost map[int]*GroupInfo
	Epoch       int
	MaxCost     int
	MaxVictory  int
	MaxCount    int
}

func (w *World) GetStatPerCost() *CostStat {
	// the stats we'll return
	stat := new(CostStat)
	stat.InfoPerCost = make(map[int]*GroupInfo)

	for cost, group := range w.perCost {
		if cost > stat.MaxCost {
			stat.MaxCost = cost
		}

		info := new(GroupInfo)
		info.GroupCount = len(group)
		// tracking the max of group count per cost
		if info.GroupCount > stat.MaxCount {
			stat.MaxCount = info.GroupCount
		}

		// we keep the best and the worst fighter in the group
		info.Winner = group[0]
		info.Loser = group[info.GroupCount-1]

		// tracking the max victory
		if info.Winner.victory > stat.MaxVictory {
			stat.MaxVictory = info.Winner.victory
		}

		stat.InfoPerCost[cost] = info
	}

	return stat
}

func (w *World) ExportLdjson(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)

	for _, fighter := range w.pool {
		enc.Encode(fighter) // LD-JSON format
	}
}
