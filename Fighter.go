package main

import "fmt"

const FIGHTING int8 = 0
const PARRY_BONUS int8 = 1

type Fighter struct {
	wounds  int
	victory int
	genome  [2]Gene
}

func (npc *Fighter) getAttack() int {
	return explodingDice(npc.getFighting())
}

func (npc *Fighter) getFighting() int {
	return npc.genome[FIGHTING].get()
}

func (target *Fighter) receiveAttack(attacker *Fighter) {
	if attacker.getAttack() >= target.getParry() {
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
	return npc.genome[FIGHTING].(*Trait).getPassiveDefense() + npc.genome[PARRY_BONUS].get()
}

func (npc *Fighter) incVictory() {
	npc.victory++
}

func (npc *Fighter) getCost() int {
	return 0
}

func buildFighter(fighting int, parryBonus int) Fighter {
	f := Fighter{}
	f.genome[FIGHTING] = &Trait{dice: fighting}
	f.genome[PARRY_BONUS] = &CappedBonus{parryBonus, 0, 2}

	return f
}

func (npc Fighter) String() string {
	return fmt.Sprint("Att:", npc.getFighting(), " ", "Dodge:", npc.genome[PARRY_BONUS].get())
}
