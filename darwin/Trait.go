package darwin

import "math/rand"

// TRAIT
type Trait struct {
	dice int
}

func (gene Trait) get() int {
	return gene.dice
}

func (gene *Trait) set(val int) {
	gene.dice = val
}

func (gene *Trait) mutate() {
	gene.dice += 4*rand.Intn(2) - 2

	if gene.dice > 12 {
		gene.dice = 12
	}
	if gene.dice < 4 {
		gene.dice = 4
	}
}

func (gene *Trait) getPassiveDefense() int {
	return gene.dice/2 + 2
}

func (gene Trait) getCost() int {
	return (gene.dice - 4) / 2
}
