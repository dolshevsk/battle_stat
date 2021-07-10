package board

type MinionJSON struct {
	Name         string
	Damage       int
	HP           int
	Type         string
	PreHitEffect string
	Deathrattles []string
}

func MinionJSONtoMinion(mJSON MinionJSON) Minion {
	return Minion{
		Name:         mJSON.Name,
		Damage:       mJSON.Damage,
		HP:           mJSON.HP,
		Type:         mJSON.Type,
		PreHitEffect: mapAttackEffects(mJSON.PreHitEffect),
		Deathrattles: mapDeathrattles(mJSON.Deathrattles),
	}
}
