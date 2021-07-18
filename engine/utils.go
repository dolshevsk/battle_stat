package engine

import (
	"battle_stat/board"
	"fmt"
)

func logCombat(boardBottom, boardTop board.Board) {
	fmt.Println("----------------------")
	fmt.Println(boardTop)
	fmt.Println(boardBottom)
	fmt.Println("----------------------")
}
