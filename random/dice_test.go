package random

import (
	"main/assert"
	"testing"
)

const enoughIteration = 1000000

func TestAverageExplodingD4(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += ExplodingDice(4)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 3.3, avg, 0.1, "Average of d4R")
}

func TestAverageExplodingD6(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += ExplodingDice(6)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 4.2, avg, 0.1, "Average of d6R")
}

func TestAverageExplodingD8(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += ExplodingDice(8)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 5.1, avg, 0.1, "Average of d8R")
}

func TestAverageExplodingD12(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += ExplodingDice(12)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 7.1, avg, 0.1, "Average of d12R")
}

func TestAverageJokerRollingD4(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += JokerRoll(4)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 5.3, avg, 0.1, "Average of joker roll d4")
}

func TestAverageJokerRollingD12(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += JokerRoll(12)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.0, avg, 0.1, "Average of joker roll d12")
}

func TestAverageRandomizeTrait(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += RandomTrait()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.0, avg, 0.1, "Average of randomized trait")
}

func TestJokerRollRateOfFire1(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += JokerRollRateOfFire(12, 1)[0] // the same as TestAverageJokerRollingD12 above
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.0, avg, 0.1, "Average of joker roll with RoF=1 for d12")
}

func TestJokerRollRateOfFire3(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += JokerRollRateOfFire(12, 3)[0]
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 11.4, avg, 0.1, "Average of joker roll with RoF=3 (3d12 + d6)keep3")
}
