package engine_test

import "battle_stat/board"

type minionToCompare struct {
	Name           string
	Damage         int
	HP             int
	IsTaunt        bool
	IsDivineShield bool
	Type           string
	IsLegendary    bool
	Tier           int8
}

func mapMinionToCompare(m board.Minion) minionToCompare {
	return minionToCompare{
		Name:           m.Name,
		Damage:         m.Damage,
		HP:             m.HP,
		IsTaunt:        m.IsTaunt,
		IsDivineShield: m.IsDivineShield,
		Type:           m.Type,
		IsLegendary:    m.IsLegendary,
		Tier:           m.Tier,
	}
}
