package main

type Fighter struct {
	fighting   int
	parryBonus int
	wounds     int
	winning    int
}

func (attacker *Fighter) attack(opponent *Fighter) {
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
