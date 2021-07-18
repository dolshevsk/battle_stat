package main

import (
	"battle_stat/data"
	"battle_stat/engine"
	"battle_stat/player"
	"log"
)

func main() {
	dataMinion, err := data.LoadData("./data/minions.json")
	if err != nil {
		log.Fatal(err)
	}

	you, err := player.NewPlayer(
		"you",
		data.MapMinions(dataMinion, "Harvest Golem"),
	)
	if err != nil {
		log.Fatal("Can't create 'you' Board", err)
	}

	opponent, err := player.NewPlayer(
		"opponent",
		data.MapMinions(dataMinion, "Glyph Guardian"),
	)
	if err != nil {
		log.Fatal("Can't create 'opponent' Board", err)
	}

	engine.RunCombat(you, opponent)
}
