package sumnseries_test

import (
	"github.com/gonzispina/gorithms/sumnseries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N=4", func(t *testing.T) {
		res := sumnseries.Solution(2)
		assert.Equal(t, int32(4), res)
	})

	t.Run("Test Solution N=10", func(t *testing.T) {
		res := sumnseries.Solution(10)
		assert.Equal(t, int32(100), res)
	})

	t.Run("Test Solution N=229137999", func(t *testing.T) {
		res := sumnseries.Solution(229137999)
		assert.Equal(t, int32(218194447), res)
	})
}
