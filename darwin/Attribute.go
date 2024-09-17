package darwin

import "math/rand"

// Attribute
type Attribute struct {
	dice int
}

func (gene Attribute) get() int {
	return gene.dice
}

func (gene *Attribute) set(val int) {
	gene.dice = val
}

func (gene *Attribute) mutate() {
	gene.dice += 4*rand.Intn(2) - 2

	if gene.dice > 12 {
		gene.dice = 12
	}
	if gene.dice < 4 {
		gene.dice = 4
	}
}

func (gene *Attribute) getPassiveDefense() int {
	return gene.dice/2 + 2
}

func (gene Attribute) getCost() int {
	return gene.dice - 4
}
