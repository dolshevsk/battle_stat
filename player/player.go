package player

import "battle_stat/board"

type Hero struct {
}

type Player struct {
	Signature string
	Hero      Hero
	Board     board.Board
}

func NewPlayer(signature string, minions []board.Minion) (Player, error) {
	playerBoard, err := board.NewBoard(minions)
	if err != nil {
		return Player{}, err
	}
	player := Player{
		Signature: signature,
		Hero:      Hero{},
		Board:     playerBoard,
	}
	return player, nil
}
