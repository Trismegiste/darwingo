package darwin

import (
	"main/assert"
	"math"
	"testing"
)

func Test_SkillCost(t *testing.T) {
	gene := new(Skill)

	gene.set(4)
	assert.AssertInt(t, 0, gene.getCost(), "d4 costs 0")
	gene.set(8)
	assert.AssertInt(t, 2, gene.getCost(), "d8 costs 2")
	gene.set(12)
	assert.AssertInt(t, 4, gene.getCost(), "d12 costs 4")
}

func (sk Skill) getTotalCost(attr int) int {
	return sk.getCost() + sk.getAdditionalCost(attr)
}

func Test_AdditionalCost_D4(t *testing.T) {
	gene := Skill{4}
	gene.set(4)

	assert.AssertInt(t, 0, gene.getAdditionalCost(4), "No additional cost for d4 with d4 attribute")
	assert.AssertInt(t, 0, gene.getAdditionalCost(12), "No additional cost for d4 with d12 attribute")
}

func Test_AdditionalCost_D12(t *testing.T) {
	gene := new(Skill)
	gene.set(12)

	assert.AssertInt(t, 0, gene.getAdditionalCost(12), "No additional cost for d12 with d12 attribute")
	assert.AssertInt(t, 1, gene.getAdditionalCost(10), "d12 with d10 cost one addtional slot")

	assert.AssertInt(t, 4, gene.getTotalCost(12), "d12 skill with d12 attribute")
	assert.AssertInt(t, 5, gene.getTotalCost(10), "d12 skill with d10 attribute")
	assert.AssertInt(t, 6, gene.getTotalCost(8), "d12 skill with d8 attribute")
	assert.AssertInt(t, 8, gene.getTotalCost(4), "d12 skill with d4 attribute")
	assert.AssertInt(t, 4, gene.getAdditionalCost(4), "Additional cost for d12 with d4 attribute")
}

func Test_SkillMutationStat(t *testing.T) {
	var countUp, countDown int = 0, 0
	for range 20000 {
		sk := Skill{8}
		sk.mutate()
		if sk.get() > 8 {
			countUp++
		}

		if sk.get() < 8 {
			countDown++
		}

		if sk.get() == 8 {
			t.Fatal("Skill should always mutate")
		}
	}

	sigma := math.Abs(float64(countDown-countUp) / 200.0)
	if sigma > 5 {
		t.Fatal("Mutation direction is biased above 5%", sigma)
	}
}

func Test_SkillEvolutionTo_D4(t *testing.T) {
	sk := Skill{8}
	for {
		sk.mutate()
		if sk.get() == 4 {
			break
		}
	}
}

func Test_SkillEvolutionTo_D12(t *testing.T) {
	sk := Skill{8}
	for {
		sk.mutate()
		if sk.get() == 12 {
			break
		}
	}
}
