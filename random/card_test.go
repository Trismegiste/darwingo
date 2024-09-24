package random

import (
	"main/assert"
	"testing"
)

func TestPickACard(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += PickCard()
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 27.5, avg, 0.1, "Average of deck")
}

func TestPickBestofTwoCards(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += PickBestCard(2)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 36.5, avg, 0.1, "Average of best between 3 cards")
}

func TestPickBestofThreeCards(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += PickBestCard(3)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 41, avg, 0.1, "Average of best between 3 cards")
}

func TestPickFirstCardAboveMaxOne(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += PickFirstCardAbove(20, 1)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 27.5, avg, 0.1, "Average of first card above 20 (max 1) = average of deck")
}

func TestPickFirstCardAboveMaxThree(t *testing.T) {
	sum := 0
	for range enoughIteration {
		sum += PickFirstCardAbove(20, 3)
	}

	var avg float64 = float64(sum) / float64(enoughIteration)
	assert.AssertFloat(t, 36.1, avg, 0.1, "Average of first card above 20 (max 3)")
}
