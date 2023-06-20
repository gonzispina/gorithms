package moretotheleft

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution true case", func(t *testing.T) {
		res := moreToTheLeft([]int{8, 6, 7, 4, 5, 1, 3, 2})
		assert.True(t, res)
	})

	t.Run("Test Solution false case", func(t *testing.T) {
		res := moreToTheLeft([]int{8, 4, 7, 6, 5, 1, 3, 2})
		assert.False(t, res)
	})
}
