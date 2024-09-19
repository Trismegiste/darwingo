package darwin

import (
	"main/assert"
	"math"
	"testing"
)

func Test_AttributeCost(t *testing.T) {
	attr := new(Attribute)

	attr.set(4)
	assert.AssertInt(t, 0, attr.getCost(), "d4 costs 0")

	attr.set(8)
	assert.AssertInt(t, 4, attr.getCost(), "d8 costs 4")

	attr.set(12)
	assert.AssertInt(t, 8, attr.getCost(), "d12 costs 8")
}

func Test_AttributeDefense(t *testing.T) {
	attr := new(Attribute)

	attr.set(4)
	assert.AssertInt(t, 4, attr.getPassiveDefense(), "Defense for d4")

	attr.set(8)
	assert.AssertInt(t, 6, attr.getPassiveDefense(), "Defense for d8")

	attr.set(12)
	assert.AssertInt(t, 8, attr.getPassiveDefense(), "Defense for d12")
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
