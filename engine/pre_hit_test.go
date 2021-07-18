package engine_test

import (
	"battle_stat/board"
	"battle_stat/engine"
	"testing"
)

func TestDoubleDamageEffect(t *testing.T) {
	you := boardFactory(t, board.Minion{
		Damage:       3,
		HP:           5,
		PreHitEffect: board.DoubleDamageEffect,
	})

	opp := boardFactory(t, board.Minion{
		Damage: 1,
		HP:     6,
	})

	expected := boardFactory(t, board.Minion{
		Damage:       6,
		HP:           4,
		PreHitEffect: board.DoubleDamageEffect,
	})

	you, opp = engine.MakeTurn(you, opp, true)
	assertEqualBoards(t, expected, you)
	assertEqualBoards(t, board.Board{}, opp)
}
