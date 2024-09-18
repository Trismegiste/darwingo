package darwin

import "math/rand"

type Strategy struct {
	current int // index of current strategy
	choices int // how many strategies are possible ?
}

func (gene Strategy) get() int {
	return gene.current
}

func (gene *Strategy) set(val int) {
	gene.current = val
}

func (gene *Strategy) mutate() {
	gene.current = rand.Intn(gene.choices)
}

func (gene Strategy) getCost() int {
	return 0
}
