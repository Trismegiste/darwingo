package random

import (
	"math"
	"testing"
)

func TestAverageExplodingD6(t *testing.T) {
	sum := 0
	iter := 1000000
	for k := 0; k < iter; k++ {
		sum += ExplodingDice(6)
	}

	var avg float64 = float64(sum) / float64(iter)
	if math.Abs(avg-4.2) > 0.1 {
		t.Fatal("Average of d6R is not around 4.2", "(", avg, ")")
	}
}
