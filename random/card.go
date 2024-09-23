package random

import "math/rand"

func PickCard() int {
	return 1 + rand.Intn(54)
}
