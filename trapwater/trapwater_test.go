package trapwater

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		res := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		assert.Equal(t, 6, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := trap([]int{4, 2, 0, 3, 2, 5})
		assert.Equal(t, 9, res)
	})

}
