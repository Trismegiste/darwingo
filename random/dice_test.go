package random

import (
	"main/assert"
	"testing"
)

const enoughIteration = 1000000

func TestAverageExplodingD4(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += ExplodingDice(4)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 3.3, avg, 0.1, "Average of d4R")
}

func TestAverageExplodingD6(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += ExplodingDice(6)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 4.2, avg, 0.1, "Average of d6R")
}

func TestAverageExplodingD8(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += ExplodingDice(8)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 5.1, avg, 0.1, "Average of d8R")
}

func TestAverageExplodingD12(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += ExplodingDice(12)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 7.1, avg, 0.1, "Average of d12R")
}

func TestAverageJokerRollingD4(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += JokerRoll(4)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 5.3, avg, 0.1, "Average of joker roll d4")
}

func TestAverageJokerRollingD12(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += JokerRoll(12)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.0, avg, 0.1, "Average of joker roll d12")
}

func TestAverageRandomizeTrait(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += RandomTrait()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 8.0, avg, 0.1, "Average of randomized trait")
}
