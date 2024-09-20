package darwin

import (
	"main/assert"
	"testing"
)

const enoughIteration = 1000000

func TestFactory(t *testing.T) {
	var f *Fighter = BuildFighter(8, 1, 6, 4, 4, 0, ATTMODE_STANDARD, 4)

	assert.AssertInt(t, 8, f.getFighting(), "Fighting skill initialisation")
	assert.AssertInt(t, 7, f.getParry(), "Parry")
	assert.AssertInt(t, 5, f.getToughness(), "Toughness")
	assert.AssertInt(t, 4, f.genome[STRENGTH].get(), "Strength")
}

func TestAttackRoll(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_SHAKEN, ATTMODE_STANDARD, 4)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getAttackRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.1, avg, 0.1, "Average of joker attack at d12")
}

func TestAttackRollWithBenny(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_ATTACK, ATTMODE_STANDARD, 4)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getAttackRoll()
		f.resetFight() // reset the number of bennies
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.8, avg, 0.1, "Average of joker attack at d12 with a benny re-roll")
}

func TestDamageRollWithCappedStrength(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_SOAK, ATTMODE_STANDARD, 4)
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
	var f *Fighter = BuildFighter(12, 0, 4, 8, 4, 0, ATTMODE_STANDARD, 4)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// therefore, we should get the average for 2d8R = 2×5.13
	assert.AssertFloat(t, 10.3, avg, 0.1, "Average of damage roll is not consistent with 2d8R")
}

func TestDamageRollWithStrength_D12(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 12, 4, 0, ATTMODE_STANDARD, 4)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// we should get the average for d8R+d12R
	assert.AssertFloat(t, 12.2, avg, 0.1, "Average of damage roll is not consistent with d8R+d12R")
}

func TestDamageRollWithStrength_D4_And_Benny(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, BENNY_TO_DAMAGE, ATTMODE_STANDARD, 4)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
		f.resetFight() // reset the number of bennies
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.5, avg, 0.1, "Average of damage roll is not consistent with 2d4R and a benny")
}

func Test_ZeroFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(4, 0, 4, 4, 4, 0, ATTMODE_STANDARD, 4)

	assert.AssertInt(t, 0, f.getCost(), "Cost should be equal to 0")
}

func Test_MiddleFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(12, 1, 6, 6, 6, 0, ATTMODE_STANDARD, 4)

	assert.AssertInt(t, (1+6)+2+2+2+2, f.getCost(), "Cost should be equal to 15")
}

func Test_HighFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(12, 1, 10, 8, 10, 0, ATTMODE_STANDARD, 4)

	assert.AssertInt(t, (3+2)+2+6+4+6, f.getCost(), "Cost should be equal to 23")
}
