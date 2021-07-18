package engine_test

import (
	"battle_stat/board"
	"github.com/stretchr/testify/require"
	"testing"
)

func boardFactory(t *testing.T, minions ...board.Minion) board.Board {
	t.Helper()
	b, err := board.NewBoard(minions)
	require.NoError(t, err)
	return b
}
