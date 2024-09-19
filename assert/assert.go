package assert

import (
	"math"
	"testing"
)

func AssertFloat(t *testing.T, expected float64, testedValue float64, delta float64, message string) {
	if math.Abs(expected-testedValue) > delta {
		t.Fatal("In", message, ": Expecting", expected, "+/-", delta, "got", testedValue)
	}
}

func AssertInt(t *testing.T, expected int, testedValue int, message string) {
	if expected != testedValue {
		t.Fatal("In", message, ": Expecting", expected, "got", testedValue)
	}
}
