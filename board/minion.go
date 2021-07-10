package board

import (
	"fmt"
)

type Minion struct {
	Name             string
	Damage           int
	HP               int
	Type             string
	PreHitEffect     func(Board, uint8) Board
	PreReceiveEffect func(Board, uint8) Board
	PostHitEffect    func(Board, uint8) Board
	Deathrattles     []Deathrattle
}

func (m Minion) String() string {
	return fmt.Sprintf("%v [%v|%v]", m.Name, m.Damage, m.HP)
}

func (m Minion) IsAlive() bool {
	return m.HP > 0
}

func (m Minion) IsDead() bool {
	return m.HP < 0
}
