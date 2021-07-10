package data

import (
	"battle_stat/board"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type MinionData map[string]board.MinionJSON

func LoadData(path string) (MinionData, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data MinionData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func MapMinions(data MinionData, names ...string) []board.Minion {
	minions := make([]board.Minion, 0, 7)
	for _, name := range names {
		minionJSON, ok := data[name]
		if !ok {
			log.Fatal(fmt.Sprintf("Can't find %v minion", name))
		}
		minions = append(minions, board.MinionJSONtoMinion(minionJSON))
	}
	return minions
}
