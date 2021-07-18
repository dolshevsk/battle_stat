package engine_test

import (
	"battle_stat/board"
	"battle_stat/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlyphGuardianVsHarvestGolemFight(t *testing.T) {
	you := boardFactory(t, board.Minion{
		Name:         "Glyph Guardian",
		Damage:       2,
		HP:           4,
		PreHitEffect: board.DoubleDamageEffect,
		Tier:         2,
	})

	opp := boardFactory(t, board.Minion{
		Name:         "Harvest Golem",
		Damage:       2,
		HP:           4,
		Deathrattles: []board.Deathrattle{board.GolemDeathrattle},
		Tier:         2,
	})

	result, damage, turn := engine.Fight(you, opp, true)
	assert.Equal(t, engine.Draw, result)
	assert.Equal(t, int8(0), damage)
	assert.Equal(t, 2, turn)
}

func TestHarvestGolemVsGlyphGuardianFight(t *testing.T) {
	you := boardFactory(t, board.Minion{
		Name:         "Harvest Golem",
		Damage:       2,
		HP:           3,
		Deathrattles: []board.Deathrattle{board.GolemDeathrattle},
		Tier:         2,
	})

	opp := boardFactory(t, board.Minion{
		Name:         "Glyph Guardian",
		Damage:       2,
		HP:           4,
		PreHitEffect: board.DoubleDamageEffect,
		Tier:         2,
	})

	result, damage, turn := engine.Fight(you, opp, true)
	assert.Equal(t, engine.Win, result)
	assert.Equal(t, int8(1), damage)
	assert.Equal(t, 2, turn)
}
