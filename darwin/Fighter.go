package darwin

import (
	"encoding/json"
	"log"
	"main/random"
	"math/rand"
)

// Indices of gene in the genome of Fighter
const (
	STRENGTH int8 = iota
	AGILITY
	VIGOR
	SPIRIT
	FIGHTING
	BENNY_STRAT
	ATTACK_MODE
	EDGE_BLOCK
	EDGE_COMBAT_REF
	EDGE_TRADEMARK_W
	EDGE_NERVE_STEEL
	EDGE_LEVEL_HEAD
	EDGE_QUICK_DRAW
)

const DEFAULT_DAMAGE_DICE = 8

// Fighter type
type Fighter struct {
	wounds       int
	victory      int
	genome       [13]Gene
	usedBenny    int
	benniesCount int
	shaken       bool
	meleeWeapon  int
}

func (f *Fighter) resetRound() {
	f.tryUnshake()
}

func (f *Fighter) tryUnshake() {
	unshake := f.rollAttr(SPIRIT) + f.genome[EDGE_COMBAT_REF].(*CappedBonus).get()*2 // adding bonus of Combat Reflexe Edge

	if unshake >= 4 {
		f.shaken = false // new rule for SWADE
		return
	}

	if f.canUseBennyStrategy(BENNY_TO_SHAKEN) {
		f.shaken = false
		f.useBenny()
	}
}

func (npc *Fighter) getAttackRoll() int {
	// if shaken, no attack
	if npc.shaken {
		return 0
	}

	att := npc.rollSkill(FIGHTING)

	// if attack has failed and the benny strategy is attack, re-roll with a benny (if any benny left)
	if att < 4 && npc.canUseBennyStrategy(BENNY_TO_ATTACK) {
		npc.useBenny()
		att = npc.rollSkill(FIGHTING)
	}

	// add bonus from wild attack
	att += npc.genome[ATTACK_MODE].(*WildAttack).getAttBonus()
	// add bonus from Trademark Weapon Edge
	att += npc.genome[EDGE_TRADEMARK_W].(*CappedBonus).get()

	return att
}

// Gets the dice for Fighting skill
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
			if target.canUseBennyStrategy(BENNY_TO_SHAKEN) {
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
		if target.canUseBennyStrategy(BENNY_TO_SOAK) {
			// use benny strategy is to soak wounds, we try
			target.useBenny()
			soak := target.rollAttr(VIGOR) / 4
			w -= soak
		}
		// if wounds are canceled by the soak roll, no wounds and no shaken
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
	ignore := npc.genome[EDGE_NERVE_STEEL].(*CappedBonus).get()
	if npc.wounds > ignore {
		return ignore - npc.wounds
	} else {
		return 0
	}
}

// test if the fighter use the given strategy for bennies and if she has a benny left
func (f *Fighter) canUseBennyStrategy(strat int) bool {
	return f.genome[BENNY_STRAT].(*Strategy).get() == strat && f.hasBenny()
}

func (npc *Fighter) getDamageRoll() int {
	roll := npc.rollDamage()

	if roll < 8 && npc.canUseBennyStrategy(BENNY_TO_DAMAGE) {
		npc.useBenny()
		roll = max(roll, npc.rollDamage())
	}

	roll += npc.genome[ATTACK_MODE].(*WildAttack).getAttBonus()

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
	parry := npc.genome[FIGHTING].(*Skill).getPassiveDefense()
	parry += npc.genome[EDGE_BLOCK].get()
	parry += npc.genome[ATTACK_MODE].(*WildAttack).getParryMalus()

	return parry
}

func (npc *Fighter) getToughness() int {
	return npc.genome[VIGOR].(*Attribute).getPassiveDefense()
}

func (f *Fighter) getInitiative() int {
	defaultDraw := f.genome[EDGE_LEVEL_HEAD].(*CappedBonus).get() + 1
	choose := random.PickBestCard(defaultDraw)

	if choose <= QUICK_DRAW_MIN_CARD && f.genome[EDGE_QUICK_DRAW].(*CappedBonus).get() == 1 {
		choose = random.PickFirstCardAbove(QUICK_DRAW_MIN_CARD, 3)
	}

	return choose
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

// Mimic the genome (we don't use a straight clone for speed, other properties are reset at the next epoch)
func (npc *Fighter) mimic(original *Fighter) {
	for idx, gene := range npc.genome {
		gene.set(original.genome[idx].get())
	}
}

// Factory
func BuildFighter(fighting int, blockEdge int, vig int, str int, agi int,
	bennyStrat int, attMode int, spi int, trademarkEdge int, combatRefEdge int,
	nerveSteel int, levelHead int, quickdraw int) *Fighter {
	f := Fighter{}
	f.genome[FIGHTING] = &Skill{fighting}
	f.genome[EDGE_BLOCK] = &CappedBonus{blockEdge, 0, 2}
	f.genome[VIGOR] = &Attribute{vig}
	f.genome[STRENGTH] = &Attribute{str}
	f.genome[AGILITY] = &Attribute{agi}
	f.genome[SPIRIT] = &Attribute{spi}
	f.genome[BENNY_STRAT] = &Strategy{bennyStrat, 4}
	f.genome[ATTACK_MODE] = &WildAttack{attMode}
	f.genome[EDGE_TRADEMARK_W] = &CappedBonus{trademarkEdge, 0, 2}
	f.genome[EDGE_COMBAT_REF] = &CappedBonus{combatRefEdge, 0, 1}
	f.genome[EDGE_NERVE_STEEL] = &CappedBonus{nerveSteel, 0, 2}
	f.genome[EDGE_LEVEL_HEAD] = &CappedBonus{levelHead, 0, 2}
	f.genome[EDGE_QUICK_DRAW] = &CappedBonus{quickdraw, 0, 1}
	f.meleeWeapon = DEFAULT_DAMAGE_DICE
	f.benniesCount = 3

	return &f
}

// Print
func (npc Fighter) String() string {
	dump, _ := json.Marshal(npc)
	return string(dump[:])
}

func (f Fighter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Victory int
		Cost    int
		STR     int
		AGI     int
		VIG     int
		SPI     int
		Fight   int
		Block   int
		TradW   int
		CmbRef  int
		NervSt  int
		LvlHd   int
		QuDrw   int
		AttMod  int
		BenStr  int
	}{
		Victory: f.victory,
		Cost:    f.getCost(),
		STR:     f.genome[STRENGTH].get(),
		AGI:     f.genome[AGILITY].get(),
		VIG:     f.genome[VIGOR].get(),
		SPI:     f.genome[SPIRIT].get(),
		Fight:   f.genome[FIGHTING].get(),
		Block:   f.genome[EDGE_BLOCK].get(),
		TradW:   f.genome[EDGE_TRADEMARK_W].get(),
		CmbRef:  f.genome[EDGE_COMBAT_REF].get(),
		NervSt:  f.genome[EDGE_NERVE_STEEL].get(),
		QuDrw:   f.genome[EDGE_QUICK_DRAW].get(),
		AttMod:  f.genome[ATTACK_MODE].get(),
		BenStr:  f.genome[BENNY_STRAT].get(),
	})
}
