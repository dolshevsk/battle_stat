package engine

import (
	"battle_stat/board"
	"battle_stat/player"
	"fmt"
	"math/rand"
	"time"
)

const (
	Win  = "WIN"
	Draw = "Draw"
	Lose = "Lose"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RunCombat(you, opp player.Player) {
	fmt.Println("Your Board:", you, "\nEnemy Board:", opp)

	var result string
	youBoard, _ := board.NewBoard(you.Minions)
	oppBoard, _ := board.NewBoard(opp.Minions)
	Fight(youBoard, oppBoard)
	fmt.Printf("Result: %v\n", result)
}

func Fight(you, opp board.Board) string {
	var turn uint
	isMyTurn := checkForMyTurn(you, opp)
	for len(you.Minions) != 0 && len(opp.Minions) != 0 {
		turn += 1
		fmt.Printf("\nTURN:%v\n", turn)
		logCombat(you, opp)
		you, opp = MakeTurn(you, opp, isMyTurn)
		logCombat(you, opp)
		isMyTurn = !isMyTurn
	}
	fmt.Printf("Battle finished in %v turns\n", turn)
	if len(you.Minions) > len(opp.Minions) {
		return Win
	} else if len(you.Minions) < len(opp.Minions) {
		return Lose
	} else {
		return Draw
	}
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

func makeAttack(bAttack, bReceive board.Board, attackerIndex, receiverIndex uint8) (board.Board, board.Board) {
	// Pre phase
	bAttack = preHit(bAttack, attackerIndex)
	bReceive = preReceive(bReceive, receiverIndex)
	// Hit phase
	fmt.Printf("%v hit %v\n", bAttack.Minions[attackerIndex], bReceive.Minions[receiverIndex])
	hit(bAttack, bReceive, attackerIndex, receiverIndex)
	// Post phase
	bAttack, bReceive = postHit(bAttack, bReceive, attackerIndex, receiverIndex)
	fmt.Printf("%v\n%v\n", bReceive, bAttack)
	return bAttack, bReceive
}

func preHit(board board.Board, attackerIndex uint8) board.Board {
	attacker := board.Minions[attackerIndex]
	if attacker.PreHitEffect == nil {
		return board
	}
	return attacker.PreHitEffect(board, attackerIndex)
}

func preReceive(board board.Board, receiverIndex uint8) board.Board {
	receiver := board.Minions[receiverIndex]
	if receiver.PreReceiveEffect == nil {
		return board
	}
	return receiver.PreHitEffect(board, receiverIndex)
}

func hit(bAttack, bReceive board.Board, attackerIndex, receiverIndex uint8) {
	attacker, receiver := &bAttack.Minions[attackerIndex], &bReceive.Minions[receiverIndex]
	receiver.HP -= attacker.Damage
	attacker.HP -= receiver.Damage
}

func postHit(bAttack, bReceive board.Board, attackerIndex, receiverIndex uint8) (board.Board, board.Board) {
	attacker := &bAttack.Minions[attackerIndex]
	if attacker.IsDead() {
		for _, deathrattle := range attacker.Deathrattles {
			bAttack, bReceive = deathrattle(bAttack, bReceive, attackerIndex)
		}
	}

	receiver := &bReceive.Minions[receiverIndex]
	if receiver.IsDead() {
		for _, deathrattle := range receiver.Deathrattles {
			bAttack, bReceive = deathrattle(bAttack, bReceive, receiverIndex)
		}
	}

	bAttack.Clean()
	bReceive.Clean()
	return bAttack, bReceive
}
