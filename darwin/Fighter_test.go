package darwin

import (
	"encoding/json"
	"main/assert"
	"math/rand"
	"testing"
)

const enoughIteration = 1000000

func TestFactory(t *testing.T) {
	var f *Fighter = BuildFighter(8, 1, 6, 4, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)

	assert.AssertInt(t, 8, f.getFighting(), "Fighting skill initialisation")
	assert.AssertInt(t, 7, f.getParry(), "Parry")
	assert.AssertInt(t, 5, f.getToughness(), "Toughness")
	assert.AssertInt(t, 4, f.genome[STRENGTH].get(), "Strength")
}

func TestAttackRoll(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_SHAKEN, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getAttackRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.1, avg, 0.1, "Average of joker attack at d12")
}

func TestAttackRollWithBenny(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_ATTACK, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getAttackRoll()
		f.resetFight() // reset the number of bennies
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.8, avg, 0.1, "Average of joker attack at d12 with a benny re-roll")
}

func TestDamageRollWithCappedStrength(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_SOAK, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// Although the damage of the weapon is d8, the roll is capped by the strength (which is d4 here)
	// therefore, we should get the average for 2d4R, therefore 2×3.35
	assert.AssertFloat(t, 6.7, avg, 0.1, "Average of damage roll is not consistent with 2d4R")
}

func TestDamageRollWithStrength_D8(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 8, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// therefore, we should get the average for 2d8R = 2×5.13
	assert.AssertFloat(t, 10.3, avg, 0.1, "Average of damage roll is not consistent with 2d8R")
}

func TestDamageRollWithStrength_D12(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 12, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// we should get the average for d8R+d12R
	assert.AssertFloat(t, 12.2, avg, 0.1, "Average of damage roll is not consistent with d8R+d12R")
}

func TestDamageRollWithStrength_D4_And_Benny(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_DAMAGE, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
		f.resetFight() // reset the number of bennies
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.5, avg, 0.1, "Average of damage roll is not consistent with 2d4R and a benny")
}

func Test_ZeroFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(4, 0, 4, 4, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)

	assert.AssertInt(t, 0, f.getCost(), "Cost should be equal to 0")
}

func Test_MiddleFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(12, 1, 6, 6, 6, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)

	assert.AssertInt(t, (1+6)+2+2+2+2, f.getCost(), "Cost should be equal to 15")
}

func Test_HighFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(12, 1, 10, 8, 10, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)

	assert.AssertInt(t, (3+2)+2+6+4+6, f.getCost(), "Cost should be equal to 23")
}

func TestDefaultInitiative(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_DAMAGE, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	sum := 0
	for range enoughIteration {
		sum += f.getInitiative()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 27.5, avg, 0.1, "Average of default Initiative card")
}

func TestDefaultInitiative_WithQuickDraw(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_DAMAGE, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 1)
	sum := 0
	for range enoughIteration {
		sum += f.getInitiative()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 37, avg, 0.1, "Average of Initiative card with Quick Draw")
}

func TestDefaultInitiative_WithLevelHead(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_DAMAGE, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 1, 0)
	sum := 0
	for range enoughIteration {
		sum += f.getInitiative()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 36.5, avg, 0.1, "Average of Initiative card with Level Head")
}

func TestVictory(t *testing.T) {
	var f *Fighter = BuildFighter(8, 1, 6, 4, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	assert.AssertInt(t, 0, f.victory, "Default value for victory")
	f.incVictory()
	assert.AssertInt(t, 1, f.victory, "Incrementing victory")
}

func TestWounds_BigCostDifference(t *testing.T) {
	var weak *Fighter = BuildFighter(4, 0, 4, 4, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0, 0)
	var strong *Fighter = BuildFighter(12, 2, 12, 12, 12, BENNY_TO_SOAK, ATTMODE_WILD, 12, 2, 2, 1, 2, 2, 1)

	var cumulWound [2]int

	for range enoughIteration {
		strong.resetFight()
		weak.resetFight()
		strong.receiveAttack(weak)
		weak.receiveAttack(strong)
		cumulWound[0] += weak.wounds
		cumulWound[1] += strong.wounds
	}

	var avgWeak float64 = float64(cumulWound[0]) / float64(enoughIteration)
	var avgStrong float64 = float64(cumulWound[1]) / float64(enoughIteration)

	assert.AssertFloat(t, 2.8, avgWeak, 0.1, "Wounds on weak")
	assert.AssertFloat(t, 0, avgStrong, 0.1, "Wounds on strong")

}

func TestWounds_Equal_But_InitiativeBias(t *testing.T) {
	var f1 *Fighter = BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 1, 0, 0, 0, 0, 0)
	var f2 *Fighter = BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 1, 0, 0, 0, 0, 0)

	var cumulWound [2]int

	for range enoughIteration {
		f2.resetFight()
		f1.resetFight()
		f2.receiveAttack(f1) // f2 is wounded first
		f1.receiveAttack(f2)
		cumulWound[0] += f1.wounds
		cumulWound[1] += f2.wounds
	}

	var avg1 float64 = float64(cumulWound[0]) / float64(enoughIteration)
	var avg2 float64 = float64(cumulWound[1]) / float64(enoughIteration)

	assert.AssertFloat(t, 0.3, avg1, 0.1, "Wounds on fighter 1")
	assert.AssertFloat(t, 0.7, avg2, 0.1, "Wounds on fighter 2")
}

func TestWounds_Equal_NoBias(t *testing.T) {
	var fighter [2]*Fighter
	var cumulWound [2]int

	fighter[0] = BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 0, 0, 0, 0, 0, 0)
	fighter[1] = BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 0, 0, 0, 0, 0, 0)

	for range enoughIteration {
		fighter[0].resetFight()
		fighter[1].resetFight()
		// randomize intiative
		if rand.Intn(2) == 0 {
			fighter[1].receiveAttack(fighter[0])
			fighter[0].receiveAttack(fighter[1])
		} else {
			fighter[0].receiveAttack(fighter[1])
			fighter[1].receiveAttack(fighter[0])
		}

		cumulWound[0] += fighter[0].wounds
		cumulWound[1] += fighter[1].wounds
	}

	var avg1 float64 = float64(cumulWound[0]) / float64(enoughIteration)
	var avg2 float64 = float64(cumulWound[1]) / float64(enoughIteration)

	assert.AssertFloat(t, 0.5, avg1, 0.1, "Wounds on fighter 1")
	assert.AssertFloat(t, 0.5, avg2, 0.1, "Wounds on fighter 2")
}

func TestUnshakeMedium(t *testing.T) {
	fighter := BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 0, 0, 0, 0, 0, 0)

	cpt := 0
	for range enoughIteration {
		fighter.shaken = true
		fighter.tryUnshake()
		if !fighter.shaken {
			cpt++
		}
	}
	var avg float64 = float64(cpt) / float64(enoughIteration)
	assert.AssertFloat(t, 0.8, avg, 0.1, "Unshaken try")
}

func TestUnshakeMedium_WithCombatRef(t *testing.T) {
	fighter := BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 0, 0, 1, 0, 0, 0)

	cpt := 0
	for range enoughIteration {
		fighter.shaken = true
		fighter.tryUnshake()
		if !fighter.shaken {
			cpt++
		}
	}
	var avg float64 = float64(cpt) / float64(enoughIteration)
	assert.AssertFloat(t, 1, avg, 0.1, "Unshaken try with edge")
}

func TestMarshallJSON(t *testing.T) {
	fighter := BuildFighter(8, 0, 8, 8, 8, 0, ATTMODE_STANDARD, 8, 0, 0, 1, 0, 0, 0)
	content, _ := json.Marshal(fighter)
	if content[0] != '{' {
		t.Fatal("JSON does not start with '{'")
	}
}
