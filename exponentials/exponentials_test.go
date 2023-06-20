package exponentials

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		res := pow(2, 16)
		assert.Equal(t, 65536, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := pow(2, 17)
		assert.Equal(t, 65536*2, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := pow(5, 5)
		assert.Equal(t, 3125, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := pow(5, 5)
		assert.Equal(t, 3125, res)
	})
}
