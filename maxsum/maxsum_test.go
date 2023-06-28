package maxsum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
		assert.Equal(t, 6, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := maxSubArray([]int{-2, -1})
		assert.Equal(t, -1, res)
	})
}
