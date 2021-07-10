package engine

import (
	"battle_stat/board"
	"fmt"
	"math/rand"
	"time"
)

const (
	Win  = "WIN"
	Draw = "Draw"
	Lose = "Lose"
)

func reverseResult(result string) string {
	if result == Win {
		return Lose
	} else if result == Lose {
		return Win
	} else {
		return Draw
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RunBattle(you, opponent game.Board) {
	fmt.Println("Your Board:", you, "\nEnemy Board:", opponent)

	var result string
	if len(you) > len(opponent) {
		result = fight(you, opponent)
	} else if len(you) < len(opponent) {
		result = reverseResult(fight(opponent, you))
	} else {
		// random choice
		r := rand.Intn(2)
		if r == 0 {
			result = fight(you, opponent)
		} else {
			result = reverseResult(fight(opponent, you))
		}
	}
	fmt.Printf("Result: %v\n", result)
}

func fight(you, opponent game.Board) string {
	var turn uint
	for len(you) != 0 && len(opponent) != 0 {
		turn += 1
		you, opponent = makeTurn(you, opponent, turn)
	}
	fmt.Printf("Battle finished in %v turns\n", turn)
	if len(you) > len(opponent) {
		return Win
	} else if len(you) < len(opponent) {
		return Lose
	} else {
		return Draw
	}
}

func makeTurn(you, opponent game.Board, turn uint) (game.Board, game.Board) {
	//var attacker, receiver board.Minion
	fmt.Printf("\nTURN:%v\n", turn)
	fmt.Println("----------------------")
	fmt.Println(opponent)
	fmt.Println(you)
	fmt.Println("----------------------")
	boardA, boardR := chooseAttackerReceiver(you, opponent, turn)
	boardA, boardR = attack(boardR, boardA, 0, 0)
	fmt.Println("----------------------")
	fmt.Println(boardR)
	fmt.Println(boardA)
	fmt.Println("----------------------")
	return boardA, boardR
}

func chooseAttackerReceiver(you, opponent game.Board, turn uint) (game.Board, game.Board) {
	var boardA, boardR game.Board
	if turn%2 == 1 {
		boardA, boardR = you, opponent
	} else {
		boardR, boardA = opponent, you
	}
	return boardA, boardR
}

func attack(boardA, boardR game.Board, attackerIndex, receiverIndex uint8) (game.Board, game.Board) {
	// Pre phase
	boardA = preHit(boardA, attackerIndex)
	boardR = preReceive(boardR, receiverIndex)
	// Hit phase
	fmt.Printf("%v hit %v\n", boardA[attackerIndex], boardR[receiverIndex])
	hit(boardA, boardR, attackerIndex, receiverIndex)
	// Post phase
	boardA, boardR = postHit(boardA, boardR, attackerIndex, receiverIndex)
	fmt.Printf("%v\n%v\n", boardR, boardA)
	return boardA, boardR
}

func preHit(board game.Board, attackerIndex uint8) game.Board {
	attacker := board[attackerIndex]
	if attacker.PreHitEffect == nil {
		return board
	}
	return attacker.PreHitEffect(board, attackerIndex)
}

func preReceive(board game.Board, receiverIndex uint8) game.Board {
	receiver := board[receiverIndex]
	if receiver.PreReceiveEffect == nil {
		return board
	}
	return receiver.PreHitEffect(board, receiverIndex)
}

func hit(board1, board2 game.Board, attackerIndex, receiverIndex uint8) {
	attacker, receiver := &board1[attackerIndex], &board2[receiverIndex]
	receiver.HP -= attacker.Damage
	attacker.HP -= receiver.Damage
}

func postHit(board1, board2 game.Board, attackerIndex, receiverIndex uint8) (game.Board, game.Board) {
	attacker := &board1[attackerIndex]
	if attacker.IsDead() {
		for _, deathrattle := range attacker.Deathrattles {
			board1, board2 = deathrattle(board1, board2, attackerIndex)
		}
	}

	receiver := &board2[receiverIndex]
	if receiver.IsDead() {
		for _, deathrattle := range receiver.Deathrattles {
			board1, board2 = deathrattle(board1, board2, receiverIndex)
		}
	}

	board1, board2 = game.Clean(board1), game.Clean(board2)

	return board1, board2
}
