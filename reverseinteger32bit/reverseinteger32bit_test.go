package reverseinteger32bit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N=312", func(t *testing.T) {
		res := reverse(312)
		assert.Equal(t, 213, res)
	})

	t.Run("Test Solution N= ", func(t *testing.T) {
		res := reverse(-312)
		assert.Equal(t, -213, res)
	})
}
