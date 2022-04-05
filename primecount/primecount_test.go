package primecount_test

import (
	"github.com/gonzispina/gorithms/primecount"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N=1", func(t *testing.T) {
		res := primecount.Solution(1)
		assert.Equal(t, int32(0), res)
	})

	t.Run("Test Solution N=6", func(t *testing.T) {
		res := primecount.Solution(6)
		assert.Equal(t, int32(2), res)
	})

	t.Run("Test Solution N=100", func(t *testing.T) {
		res := primecount.Solution(100)
		assert.Equal(t, int32(3), res)
	})
}
