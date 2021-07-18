package player

import "battle_stat/board"

type Hero struct {
}

type Player struct {
	Signature string
	Hero      Hero
	Minions   []board.Minion
}

func NewPlayer(signature string, minions []board.Minion) (Player, error) {
	player := Player{
		Signature: signature,
		Hero:      Hero{},
		Minions:   minions,
	}
	return player, nil
}
