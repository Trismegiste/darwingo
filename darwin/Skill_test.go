package darwin

import "testing"

func Test_Cost(t *testing.T) {
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

func Test_AdditionalCost(t *testing.T) {
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
