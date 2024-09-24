package darwin

import (
	"main/assert"
	"testing"
)

func TestOneEpoch(t *testing.T) {
	sut := new(World)
	var weak *Fighter = BuildFighter(4, 0, 4, 4, 4, 0, ATTMODE_STANDARD, 4, 0, 0, 0, 0, 0)
	var strong *Fighter = BuildFighter(12, 2, 12, 12, 12, BENNY_TO_SOAK, ATTMODE_WILD, 12, 2, 1, 2, 2, 1)
	sut.pool = []*Fighter{weak, strong}

	sut.RunEpoch(30)
	assert.AssertInt(t, 1, len(sut.perCost[0]), "Stat on cost 0")
	assert.AssertInt(t, 1, len(sut.perCost[strong.getCost()]), "Stat on cost 56")
	assert.AssertInt(t, 1, strong.victory, "Strong wins")
	assert.AssertInt(t, 0, weak.victory, "Weak loses")

	stat := sut.GetStatPerCost()
	assert.AssertInt(t, 56, stat.MaxCost, "Max cost")
	assert.AssertInt(t, 1, stat.MaxCount, "1 per cost max")
	assert.AssertInt(t, 1, stat.MaxVictory, "1 victory max")
}
