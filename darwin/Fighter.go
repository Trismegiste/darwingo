package darwin

import (
	"encoding/json"
	"fmt"
	"log"
	"main/random"
	"math/rand"
)

// Indices of gene in the genome of Fighter
const (
	FIGHTING int8 = iota
	BLOCK
	VIGOR
	STRENGTH
	AGILITY
	BENNY_STRAT
)

const (
	BENNY_TO_ATTACK int = iota
	BENNY_TO_SOAK
	BENNY_TO_SHAKEN
	BENNY_TO_DAMAGE
)

const DEFAULT_DAMAGE_DICE = 8

// Fighter type
type Fighter struct {
	wounds       int
	victory      int
	genome       [6]Gene
	usedBenny    int
	benniesCount int
	shaken       bool
	meleeWeapon  int
}

func (npc *Fighter) getAttackRoll() int {
	att := npc.rollSkill(FIGHTING)
	if att < 4 && npc.genome[BENNY_STRAT].get() == BENNY_TO_ATTACK && npc.hasBenny() {
		npc.useBenny()
		att = npc.rollSkill(FIGHTING)
	}

	return att
}

func (npc *Fighter) getFighting() int {
	return npc.genome[FIGHTING].get()
}

func (target *Fighter) receiveAttack(attacker *Fighter) {
	delta := attacker.getAttackRoll() - target.getParry()
	if delta >= 0 {
		// hit ! Calculate the damage
		damage := attacker.getDamageRoll()
		// Raise ?
		if delta >= 4 {
			damage += random.ExplodingDice(6)
		}
		target.receiveDamage(damage)
	}
}

func (target *Fighter) receiveDamage(damage int) {
	// compare damage and toughness
	if damage >= target.getToughness() {
		injuries := (damage - target.getToughness()) / 4
		target.addWounds(injuries)
	}
}

func (npc *Fighter) hasBenny() bool {
	return npc.usedBenny < npc.benniesCount
}

func (npc *Fighter) useBenny() {
	if npc.usedBenny == npc.benniesCount {
		log.Fatal("No benny left")
	}
	npc.usedBenny++
}

func (target *Fighter) addWounds(w int) {
	if w == 0 {
		// new shaken condition :
		if target.shaken {
			// if already shaken, unshake before getting a new shaken by using a benny
			if (target.genome[BENNY_STRAT].get() == BENNY_TO_SHAKEN) && target.hasBenny() {
				target.useBenny()
			} else {
				target.wounds++ // 2 shaken = 1 wound
			}
		} else {
			// not already shaken but damage below 1 wound, get shaken condition
			target.shaken = true
		}
	} else {
		// receiving 1 or more wounds
		if (target.genome[BENNY_STRAT].get() == BENNY_TO_SOAK) && target.hasBenny() {
			// use benny strategy is to soak wounds, we try
			target.useBenny()
			soak := target.rollAttr(VIGOR) / 4
			w -= soak
		}
		if w > 0 {
			target.wounds += w
			target.shaken = true
		}
	}
}

func (npc *Fighter) rollAttr(idxGenome int8) int {
	// transcast to Attribute is not mandatory but it prevents to use a CappedBonus or a Strategy by mistake
	return random.JokerRoll(npc.genome[idxGenome].(*Attribute).get()) + npc.getWoundsPenalty()
}

func (npc *Fighter) rollSkill(idxGenome int8) int {
	// transcast to Skill is not mandatory but it prevents to use a CappedBonus or a Strategy by mistake
	return random.JokerRoll(npc.genome[idxGenome].(*Skill).get()) + npc.getWoundsPenalty()
}

func (npc *Fighter) getWoundsPenalty() int {
	return -npc.wounds
}

func (npc *Fighter) getDamageRoll() int {
	roll := npc.rollDamage()
	if roll < 8 && npc.genome[BENNY_STRAT].(*Strategy).get() == BENNY_TO_DAMAGE && npc.hasBenny() {
		npc.useBenny()
		roll = max(roll, npc.rollDamage())
	}

	return roll
}

func (npc *Fighter) rollDamage() int {
	str := npc.genome[STRENGTH].get()
	damageDice := npc.meleeWeapon
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
	npc.usedBenny = 0
	npc.shaken = false
}

func (npc *Fighter) resetEpoch() {
	npc.victory = 0
}

func (npc *Fighter) getParry() int {
	return npc.genome[FIGHTING].(*Skill).getPassiveDefense() + npc.genome[BLOCK].get()
}

func (npc *Fighter) getToughness() int {
	return npc.genome[VIGOR].(*Attribute).getPassiveDefense()
}

func (npc *Fighter) incVictory() {
	npc.victory++
}

func (npc *Fighter) getCost() int {
	sum := 0
	for _, gene := range npc.genome {
		sum += gene.getCost()
	}
	sum += npc.genome[FIGHTING].(*Skill).getAdditionalCost(npc.genome[AGILITY].get())

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
func BuildFighter(fighting int, blockEdge int, vig int, str int, agi int, bennyStrat int) *Fighter {
	f := Fighter{}
	f.genome[FIGHTING] = &Skill{fighting}
	f.genome[BLOCK] = &CappedBonus{blockEdge, 0, 2}
	f.genome[VIGOR] = &Attribute{vig}
	f.genome[STRENGTH] = &Attribute{str}
	f.genome[AGILITY] = &Attribute{agi}
	f.genome[BENNY_STRAT] = &Strategy{bennyStrat, 4}
	f.meleeWeapon = DEFAULT_DAMAGE_DICE
	f.benniesCount = 3

	return &f
}

// Print
func (npc Fighter) String() string {
	return fmt.Sprint("Att:", npc.getFighting(), " ",
		"VIG:", npc.genome[VIGOR].get(), " ",
		"STR:", npc.genome[STRENGTH].get(), " ",
		"AGI:", npc.genome[AGILITY].get(), " ",
		"Block:", npc.genome[BLOCK].get(), " ",
		"BS:", npc.genome[BENNY_STRAT].get(), " ",
		"Cost:", npc.getCost(), " ",
		"Win:", npc.victory)
}

func (f Fighter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Detail  string
		Victory int
	}{
		Detail:  f.String(),
		Victory: f.victory,
	})
}
