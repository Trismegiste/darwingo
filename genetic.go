package main

import "math/rand"

type Gene interface {
	get() int
	mutate()
}

// TRAIT
type Trait struct {
	dice int
}

func (gene Trait) get() int {
	return gene.dice
}
func (gene *Trait) mutate() {
	gene.dice += rand.Intn(5) - 2

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

// BONUS
type CappedBonus struct {
	value int
	min   int
	max   int
}

func (gene CappedBonus) get() int {
	return gene.value
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
