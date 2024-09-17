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
const STRENGTH int8 = 3

const DEFAULT_DAMAGE_DICE = 8

// Fighter type
type Fighter struct {
	wounds  int
	victory int
	genome  [4]Gene
}

func (npc *Fighter) getAttackRoll() int {
	return random.ExplodingDice(npc.getFighting())
}

func (npc *Fighter) getFighting() int {
	return npc.genome[FIGHTING].get()
}

func (target *Fighter) receiveAttack(attacker *Fighter) {
	delta := attacker.getAttackRoll() - target.getParry()
	if delta >= 0 {
		// hit ! Calculate the damage
		damage := attacker.getDamageRoll()
		if delta >= 4 {
			damage += random.ExplodingDice(6)
		}

		// Effect of damage
		injuries := (damage - target.getToughness()) / 4
		if injuries > 0 {
			target.wounds += injuries
		}
	}
}

func (npc *Fighter) getDamageRoll() int {
	str := npc.genome[STRENGTH].get()
	damageDice := DEFAULT_DAMAGE_DICE
	if damageDice > str {
		damageDice = str
	}

	return random.ExplodingDice(str) + random.ExplodingDice(damageDice)
}

func (npc *Fighter) isDead() bool {
	return npc.wounds > 3
}

func (npc *Fighter) resetFight() {
	npc.wounds = 0
}

func (npc *Fighter) resetEpoch() {
	npc.victory = 0
}

func (npc *Fighter) getParry() int {
	return npc.genome[FIGHTING].(*Trait).getPassiveDefense() + npc.genome[BLOCK].get()
}

func (npc *Fighter) getToughness() int {
	return npc.genome[VIGOR].(*Trait).getPassiveDefense()
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

func (npc *Fighter) mimic(original *Fighter) {
	for idx, gene := range npc.genome {
		gene.set(original.genome[idx].get())
	}
}

// Factory
func BuildFighter(fighting int, blockEdge int, vig int, str int) *Fighter {
	f := Fighter{}
	f.genome[FIGHTING] = &Trait{fighting}
	f.genome[BLOCK] = &CappedBonus{blockEdge, 0, 2}
	f.genome[VIGOR] = &Trait{vig}
	f.genome[STRENGTH] = &Trait{str}

	return &f
}

// Print
func (npc Fighter) String() string {
	return fmt.Sprint("Att:", npc.getFighting(), " ",
		"VIG:", npc.genome[VIGOR].get(), " ",
		"STR:", npc.genome[STRENGTH].get(), " ",
		"Block:", npc.genome[BLOCK].get(), " ",
		"Cost:", npc.getCost(), " ",
		"Win:", npc.victory)
}
