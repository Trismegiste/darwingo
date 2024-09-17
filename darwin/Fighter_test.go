package darwin

import "testing"

const enoughIteration = 1000000

func TestFactory(t *testing.T) {
	var f *Fighter = BuildFighter(8, 1, 6, 4)

	if f.getFighting() != 8 {
		t.Fatal("Fighting skill is not initialized to 8")
	}

	if f.getParry() != 7 {
		t.Fatal("Parry should be equal to 8/2+2 + 1")
	}

	if f.getToughness() != 5 {
		t.Fatal("Toughness should be equal to 6/2+2")
	}

	if f.genome[STRENGTH].get() != 4 {
		t.Fatal("Strength should be equal to 4")
	}
}
