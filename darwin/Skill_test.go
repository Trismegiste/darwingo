package darwin

import (
	"math"
	"testing"
)

func Test_SkillCost(t *testing.T) {
	gene := new(Skill)

	gene.set(8)
	if gene.getCost() != 2 {
		t.Fatal("d8 to d4 is 2 slots")
	}

	gene.set(12)
	if gene.getCost() != 4 {
		t.Fatal("d12 to d4 is 4 slots")
	}
}

func (sk Skill) getTotalCost(attr int) int {
	return sk.getCost() + sk.getAdditionalCost(attr)
}

func Test_AdditionalCost_D4(t *testing.T) {
	gene := Skill{4}

	if gene.getCost() != 0 {
		t.Fatal("d4 Skill costs 0")
	}

	if gene.getAdditionalCost(4) != 0 {
		t.Fatal("No additional cost for a d4 Skill to d4 when Attribute is equal to d4")
	}

	if gene.getAdditionalCost(12) != 0 {
		t.Fatal("No additional cost for a d4 Skill to d4 when Attribute is equal to d12")
	}

}

func Test_AdditionalCost_D12(t *testing.T) {
	gene := new(Skill)
	gene.set(12)

	if gene.getAdditionalCost(12) != 0 {
		t.Fatal("No additional cost for a d12 Skill to d4 when Attribute is equal to d12")
	}

	if gene.getTotalCost(12) != 4 {
		t.Fatal("d12 Skill to d4 is 4 slots if Attribute is equal to d12")
	}

	if gene.getTotalCost(10) != 5 {
		t.Fatal("d12 Skill to d4 is 5 slots if Attribute is equal to d10")
	}

	if gene.getTotalCost(8) != 6 {
		t.Fatal("d12 Skill to d4 is 7 slots if Attribute is equal to d8")
	}

	if gene.getTotalCost(4) != 8 {
		t.Fatal("d12 Skill to d4 is 8 slots if Attribute is equal to d4")
	}

	if gene.getAdditionalCost(4) != 4 {
		t.Fatal("4 additional slots for a d12 Skill to d4 when Attribute is equal to d4")
	}
}

func Test_SkillMutation(t *testing.T) {
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
