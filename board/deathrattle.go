package board

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Deathrattle func(Board, Board, int8) (Board, Board)

var deathrattlesMap = map[string]Deathrattle{
	"GolemDeathrattle": GolemDeathrattle,
}

func mapDeathrattles(deathrattleNames []string) []Deathrattle {
	deathrattles := make([]Deathrattle, len(deathrattleNames))
	for i, name := range deathrattleNames {
		deathrattle, ok := deathrattlesMap[name]
		if !ok {
			log.Fatal(fmt.Sprintf("Can't find %v deathrattle", name))
		}
		deathrattles[i] = deathrattle
	}
	return deathrattles
}

func GolemDeathrattle(deathrattled, opposite Board, deathIndex int8) (Board, Board) {
	golem := Minion{
		Name:   "Damaged Golem",
		Damage: 2,
		HP:     1,
		Type:   Mech,
		Tier:   1,
	}

	deathrattled.Insert(deathIndex, golem)
	return deathrattled, opposite
}
