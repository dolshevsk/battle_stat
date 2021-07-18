package board

import (
	"errors"
	"fmt"
)

type Board struct {
	Minions       []Minion
	AttackPointer int8
}

func NewBoard(minions []Minion) (Board, error) {
	if len(minions) > 7 {
		return Board{}, errors.New(fmt.Sprintf("Can't be more minions than 7, got %v", len(minions)))
	}
	board := Board{
		Minions: minions,
	}
	return board, nil
}

func (b *Board) Clean() {
	cleanedBoard := make([]Minion, 0, 7)
	for _, minion := range b.Minions {
		if minion.IsAlive() {
			cleanedBoard = append(cleanedBoard, minion)
		}
	}
	b.Minions = cleanedBoard
}

func (b *Board) Remove(index int8) {
	minions := append(b.Minions[:index], b.Minions[index+1:]...)
	b.Minions = minions
}

func (b *Board) Insert(index int8, minion Minion) {
	minions := append(b.Minions[:index+1], b.Minions[index:]...)
	minions[index] = minion
	b.Minions = minions
}
