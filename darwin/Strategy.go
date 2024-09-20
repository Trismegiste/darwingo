package darwin

import "math/rand"

const (
	BENNY_TO_ATTACK int = iota
	BENNY_TO_SOAK
	BENNY_TO_SHAKEN
	BENNY_TO_DAMAGE
)

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
