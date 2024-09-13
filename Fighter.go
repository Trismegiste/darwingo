package main

type Fighter struct {
	fighting   int
	parryBonus int
	wounds     int
	winning    int
	shaken     bool
}

func (attacker *Fighter) attack(opponent *Fighter) {
	if attacker.shaken {
		return
	}

	if explodingDice(attacker.fighting) >= opponent.getParry() {
		opponent.wounds++
	}
}

func (npc *Fighter) isDead() bool {
	return npc.wounds > 3
}

func (npc *Fighter) reset() {
	npc.wounds = 0
}

func (npc *Fighter) getParry() int {
	return npc.fighting/2 + 2 + npc.parryBonus
}

func (npc *Fighter) incVictory() {
	npc.winning++
}
