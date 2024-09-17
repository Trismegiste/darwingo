package darwin

import "math/rand"

// Skill
type Skill struct {
	dice int
}

func (gene Skill) get() int {
	return gene.dice
}

func (gene *Skill) set(val int) {
	gene.dice = val
}

func (gene *Skill) mutate() {
	gene.dice += 4*rand.Intn(2) - 2

	if gene.dice > 12 {
		gene.dice = 12
	}
	if gene.dice < 4 {
		gene.dice = 4
	}
}

func (gene *Skill) getPassiveDefense() int {
	return gene.dice/2 + 2
}

func (gene Skill) getCost() int {
	return (gene.dice - 4) / 2
}

func (gene Skill) getAdditionalCost(attr int) int {
	if gene.dice > attr {
		return (gene.dice - attr) / 2
	}

	return 0
}
