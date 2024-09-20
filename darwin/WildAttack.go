package darwin

import "math/rand"

const (
	ATTMODE_STANDARD = 0
	ATTMODE_WILD     = 1
)

type WildAttack struct {
	current int // see const above for meaning
}

func (gene WildAttack) get() int {
	return gene.current
}

func (gene *WildAttack) set(val int) {
	gene.current = val
}

func (gene *WildAttack) mutate() {
	gene.current = rand.Intn(2)
}

func (gene WildAttack) getCost() int {
	return 0
}

func (gene WildAttack) getAttBonus() int {
	if gene.current == ATTMODE_WILD {
		return 2
	} else {
		return 0
	}
}

func (gene WildAttack) getParryMalus() int {
	return -gene.getAttBonus()
}
