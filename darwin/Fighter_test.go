package darwin

import (
	"math"
	"testing"
)

const enoughIteration = 1000000

func TestFactory(t *testing.T) {
	var f *Fighter = BuildFighter(8, 1, 6, 4, 4, 0)

	if f.getFighting() != 8 {
		t.Fatal("Fighting skill is not initialized to 8")
	}

	if f.getParry() != 7 {
		t.Fatal("Parry should be equal to 8/2+2 + 1")
	}

	if f.getToughness() != 5 {
		t.Fatal("Toughness should be equal to 6/2+2")
	}

	if f.genome[STRENGTH].get() != 4 {
		t.Fatal("Strength should be equal to 4")
	}
}

func TestAttackRoll(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getAttackRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	if math.Abs(avg-8.0) > 0.1 {
		t.Fatal("Average of joker attack at d12 is not around 8.0", "(", avg, ")")
	}
}

func TestDamageRollWithCappedStrength(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 4, 4, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// Although the damage of the weapon is d8, the roll is capped by the strength (which is d4 here)
	// therefore, we should get the average for 2d4R
	if math.Abs(avg-6.7) > 0.1 {
		t.Fatal("Average of damage roll is not consistent with 2d4R around 2*3.3", "(", avg, ")")
	}
}

func TestDamageRollWithStrength_D8(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 8, 4, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// therefore, we should get the average for 2d8R
	if math.Abs(avg-10.3) > 0.1 {
		t.Fatal("Average of damage roll is not consistent with 2d8R around 2*5.13", "(", avg, ")")
	}
}

func TestDamageRollWithStrength_D12(t *testing.T) {
	var f *Fighter = BuildFighter(12, 0, 4, 12, 4, 0)
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += f.getDamageRoll()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	// we should get the average for d8R+d12R
	if math.Abs(avg-12.2) > 0.1 {
		t.Fatal("Average of damage roll is not consistent with d8R+d12R around 5.1+7.1", "(", avg, ")")
	}
}

func Test_ZeroFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(4, 0, 4, 4, 4, 0)

	if f.getCost() != 0 {
		t.Fatal("Cost should be equal to 0")
	}
}

func Test_MiddleFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(12, 1, 6, 6, 6, 0)

	if f.getCost() != (1+6)+2+2+2+2 {
		t.Fatal("Cost should be equal to 15", f.getCost())
	}
}

func Test_HighFighterCost(t *testing.T) {
	var f *Fighter = BuildFighter(12, 1, 10, 8, 10, 0)

	if f.getCost() != (3+2)+2+6+4+6 {
		t.Fatal("Cost should be equal to 23")
	}
}
