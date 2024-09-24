package darwin

import "math/rand"

const QUICK_DRAW_MIN_CARD = 20

// BONUS
type CappedBonus struct {
	value int
	min   int
	max   int
}

func (gene CappedBonus) get() int {
	return gene.value
}

func (gene *CappedBonus) set(val int) {
	gene.value = val
}

func (gene *CappedBonus) mutate() {
	gene.value += 2*rand.Intn(2) - 1

	if gene.value > gene.max {
		gene.value = gene.max
	}
	if gene.value < gene.min {
		gene.value = gene.min
	}
}

func (gene CappedBonus) getCost() int {
	return 2 * gene.value
}
