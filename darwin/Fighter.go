package darwin

import (
	"fmt"
	"main/random"
	"math/rand"
)

// Indices of gene in the genome of Fighter
const FIGHTING int8 = 0
const BLOCK int8 = 1
const VIGOR int8 = 2

// Fighter type
type Fighter struct {
	wounds  int
	victory int
	genome  [3]Gene
}

func (npc *Fighter) getAttackRoll() int {
	return random.ExplodingDice(npc.getFighting())
}

func (npc *Fighter) getFighting() int {
	return npc.genome[FIGHTING].get()
}

func (target *Fighter) receiveAttack(attacker *Fighter) {
	if attacker.getAttackRoll() >= target.getParry() {
		target.wounds++
	}
}

func (npc *Fighter) isDead() bool {
	return npc.wounds > 3
}

func (npc *Fighter) reset() {
	npc.wounds = 0
}

func (npc *Fighter) getParry() int {
	return npc.genome[FIGHTING].(*Trait).getPassiveDefense() + npc.genome[BLOCK].get()
}

func (npc *Fighter) incVictory() {
	npc.victory++
}

func (npc *Fighter) getCost() int {
	sum := 0
	for _, gene := range npc.genome {
		sum += gene.getCost()
	}

	return sum
}

func (npc *Fighter) mutate() {
	pick := npc.genome[rand.Intn(len(npc.genome))]
	pick.mutate()
}

// Factory
func BuildFighter(fighting int, parryBonus int, vigor int) Fighter {
	f := Fighter{}
	f.genome[FIGHTING] = &Trait{fighting}
	f.genome[BLOCK] = &CappedBonus{parryBonus, 0, 2}
	f.genome[VIGOR] = &Trait{vigor}

	return f
}

// Print
func (npc Fighter) String() string {
	return fmt.Sprint("Att:", npc.getFighting(), " ",
		"Vig:", npc.genome[VIGOR].get(), " ",
		"Dodge:", npc.genome[BLOCK].get(), " ",
		"Cost:", npc.getCost())
}
