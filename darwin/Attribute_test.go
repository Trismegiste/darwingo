package darwin

import (
	"math"
	"testing"
)

func Test_AttributeCost(t *testing.T) {
	attr := new(Attribute)

	attr.set(4)
	if attr.getCost() != 0 {
		t.Fatal("Cost should be equal to 0")
	}
	attr.set(8)
	if attr.getCost() != 4 {
		t.Fatal("Cost should be equal to 4")
	}

	attr.set(12)
	if attr.getCost() != 8 {
		t.Fatal("Cost should be equal to 8")
	}
}

func Test_AttributeDefense(t *testing.T) {
	attr := new(Attribute)

	attr.set(4)
	if attr.getPassiveDefense() != 4 {
		t.Fatal("Defense should be equal to 4")
	}
	attr.set(8)
	if attr.getPassiveDefense() != 6 {
		t.Fatal("Cost should be equal to 6")
	}

	attr.set(12)
	if attr.getPassiveDefense() != 8 {
		t.Fatal("Cost should be equal to 8")
	}
}

func Test_AttributeMutationStat(t *testing.T) {
	var countUp, countDown int = 0, 0
	for range 20000 {
		attr := Attribute{8}
		attr.mutate()
		if attr.get() > 8 {
			countUp++
		}

		if attr.get() < 8 {
			countDown++
		}

		if attr.get() == 8 {
			t.Fatal("Attribute should always mutate")
		}
	}

	sigma := math.Abs(float64(countDown-countUp) / 200.0)
	if sigma > 5 {
		t.Fatal("Mutation direction is biased above 5%", sigma)
	}
}

func Test_AttributeEvolutionTo_D4(t *testing.T) {
	attr := Attribute{8}
	for {
		attr.mutate()
		if attr.get() == 4 {
			break
		}
	}
}

func Test_AttributeEvolutionTo_D12(t *testing.T) {
	attr := Attribute{8}
	for {
		attr.mutate()
		if attr.get() == 12 {
			break
		}
	}
}
