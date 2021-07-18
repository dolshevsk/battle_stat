package board

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type AttackEffect func(Board, int8) Board

var attackEffectsMap = map[string]AttackEffect{
	"DoubleDamageEffect": DoubleDamageEffect,
}

func mapAttackEffects(name string) AttackEffect {
	if name == "" {
		return nil
	}
	effect, ok := attackEffectsMap[name]
	if !ok {
		log.Fatal(fmt.Sprintf("Can't find %v attackEffect", name))
	}
	return effect
}

func DoubleDamageEffect(board Board, index int8) Board {
	minion := &board.Minions[index]
	minion.Damage *= 2
	return board
}
