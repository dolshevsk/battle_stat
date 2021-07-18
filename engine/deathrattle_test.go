package engine_test

import (
	"battle_stat/board"
	"battle_stat/engine"
	"testing"
)

func TestGolemDeathrattle(t *testing.T) {
	you := boardFactory(t, board.Minion{
		Damage:       2,
		HP:           3,
		Deathrattles: []board.Deathrattle{board.GolemDeathrattle},
	})

	opp := boardFactory(t, board.Minion{
		Damage: 5,
		HP:     6,
	})

	expected := boardFactory(t, board.Minion{
		Name:   "Damaged Golem",
		Damage: 2,
		HP:     1,
		Type:   board.Mech,
		Tier:   1,
	})

	you, opp = engine.MakeTurn(you, opp, true)
	assertEqualBoards(t, expected, you)
}
