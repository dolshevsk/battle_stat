package engine_test

import (
	"battle_stat/board"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func assertEqualBoards(t *testing.T, expected, actual board.Board) bool {
	t.Helper()
	require.Equal(t, len(expected.Minions), len(actual.Minions))

	var e, a []minionToCompare
	for i := range actual.Minions {
		e = append(e, mapMinionToCompare(expected.Minions[i]))
		a = append(a, mapMinionToCompare(actual.Minions[i]))
	}
	return assert.Equal(t, e, a)
}
