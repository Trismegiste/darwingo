package random

import (
	"math"
	"testing"
)

const enoughIteration = 1000000

func TestAverageExplodingD6(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += ExplodingDice(6)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	if math.Abs(avg-4.2) > 0.1 {
		t.Fatal("Average of d6R is not around 4.2", "(", avg, ")")
	}
}

func TestAverageJokerRollingD4(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += JokerRoll(4)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	if math.Abs(avg-5.3) > 0.1 {
		t.Fatal("Average of joker roll d4 is not around 5.3", "(", avg, ")")
	}
}

func TestAverageJokerRollingD12(t *testing.T) {
	sum := 0
	for k := 0; k < enoughIteration; k++ {
		sum += JokerRoll(12)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	if math.Abs(avg-8.0) > 0.1 {
		t.Fatal("Average of joker roll d12 is not around 8.0", "(", avg, ")")
	}
}
