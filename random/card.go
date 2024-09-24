package random

import "math/rand"

func PickCard() int {
	return 1 + rand.Intn(54)
}

func PickBestCard(n int) int {
	best := 0
	for range n {
		draw := PickCard()
		if draw > best {
			best = draw
		}
	}

	return best
}

func PickFirstCardAbove(maxExcludedCard int, maxDraw int) int {
	draw := 0

	for range maxDraw {
		draw = PickCard()
		if draw > maxExcludedCard {
			return draw
		}
	}

	return draw
}
