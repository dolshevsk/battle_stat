package board

import (
	"fmt"
)

type Minion struct {
	Name             string
	Damage           int
	HP               int
	IsTaunt          bool
	IsDivineShield   bool
	Type             string
	IsLegendary      bool
	Tier             int8
	PreHitEffect     func(Board, int8) Board
	PreReceiveEffect func(Board, int8) Board
	PostHitEffect    func(Board, int8) Board
	Deathrattles     []Deathrattle
}

func (m Minion) String() string {
	return fmt.Sprintf("%v [%v|%v]", m.Name, m.Damage, m.HP)
}

func (m Minion) IsAlive() bool {
	return m.HP > 0
}

func (m Minion) IsDead() bool {
	return !m.IsAlive()
}
