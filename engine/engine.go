package engine

import (
	"battle_stat/board"
	"battle_stat/player"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RunCombat(you, opp player.Player) {
	fmt.Println("Your Board:", you, "\nEnemy Board:", opp)

	youBoard, _ := board.NewBoard(you.Minions)
	oppBoard, _ := board.NewBoard(opp.Minions)
	result, damage, turn := chooseWhoFirstAndFight(youBoard, oppBoard)
	fmt.Printf("Combat finished with result: %v; damage: %v; turns: %v.\n", result, damage, turn)
}

func chooseWhoFirstAndFight(you, opp board.Board) (result string, damage int8, turn int) {
	isMyTurn := checkForMyTurn(you, opp)
	return Fight(you, opp, isMyTurn)
}

func Fight(you, opp board.Board, isMyTurn bool) (result string, damage int8, turn int) {
	for len(you.Minions) != 0 && len(opp.Minions) != 0 {
		turn += 1
		fmt.Printf("\nTURN:%v\n", turn)
		logCombat(you, opp)
		you, opp = MakeTurn(you, opp, isMyTurn)
		logCombat(you, opp)
		isMyTurn = !isMyTurn
	}
	if len(you.Minions) > len(opp.Minions) {
		result = Win
		damage = calculateBoardDamage(you)
	} else if len(you.Minions) < len(opp.Minions) {
		result = Lose
		damage = calculateBoardDamage(opp)
	} else {
		result = Draw
		damage = 0
	}
	return result, damage, turn
}

func calculateBoardDamage(board board.Board) (damage int8) {
	for _, m := range board.Minions {
		damage += m.Tier
	}
	return damage
}

func checkForMyTurn(you, opp board.Board) bool {
	if len(you.Minions) > len(opp.Minions) {
		return true
	} else if len(you.Minions) < len(opp.Minions) {
		return false
	} else {
		// random choice
		r := rand.Intn(2)
		if r == 0 {
			return true
		} else {
			return false
		}
	}
}

func MakeTurn(you, opp board.Board, isMyTurn bool) (board.Board, board.Board) {
	//var attacker, receiver board.Minion
	if isMyTurn {
		you, opp = makeAttack(you, opp, you.AttackPointer, 0)
	} else {
		opp, you = makeAttack(opp, you, opp.AttackPointer, 0)
	}
	return you, opp
}

func makeAttack(bAttack, bReceive board.Board, attackerIndex, receiverIndex int8) (board.Board, board.Board) {
	// Pre phase
	bAttack = preHit(bAttack, attackerIndex)
	bReceive = preReceive(bReceive, receiverIndex)
	// Hit phase
	fmt.Printf("%v hit %v\n", bAttack.Minions[attackerIndex], bReceive.Minions[receiverIndex])
	hit(bAttack, bReceive, attackerIndex, receiverIndex)
	// Post phase
	bAttack, bReceive = postHit(bAttack, bReceive, attackerIndex, receiverIndex)
	return bAttack, bReceive
}

func preHit(board board.Board, attackerIndex int8) board.Board {
	attacker := board.Minions[attackerIndex]
	if attacker.PreHitEffect == nil {
		return board
	}
	return attacker.PreHitEffect(board, attackerIndex)
}

func preReceive(board board.Board, receiverIndex int8) board.Board {
	receiver := board.Minions[receiverIndex]
	if receiver.PreReceiveEffect == nil {
		return board
	}
	return receiver.PreHitEffect(board, receiverIndex)
}

func hit(bAttack, bReceive board.Board, attackerIndex, receiverIndex int8) {
	attacker, receiver := &bAttack.Minions[attackerIndex], &bReceive.Minions[receiverIndex]
	receiver.HP -= attacker.Damage
	attacker.HP -= receiver.Damage
}

func postHit(bAttack, bReceive board.Board, attackerIndex, receiverIndex int8) (board.Board, board.Board) {
	attacker := &bAttack.Minions[attackerIndex]
	if attacker.IsDead() {
		for _, deathrattle := range attacker.Deathrattles {
			bAttack, bReceive = deathrattle(bAttack, bReceive, attackerIndex)
		}
	}

	receiver := &bReceive.Minions[receiverIndex]
	if receiver.IsDead() {
		for _, deathrattle := range receiver.Deathrattles {
			bReceive, bAttack = deathrattle(bReceive, bAttack, receiverIndex)
		}
	}

	bAttack.Clean()
	bReceive.Clean()
	return bAttack, bReceive
}
