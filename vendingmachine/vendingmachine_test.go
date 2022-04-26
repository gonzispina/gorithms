package vendingmachine_test

import (
	"github.com/gonzispina/gorithms/vendingmachine"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution M=3.14, P=1.99", func(t *testing.T) {
		res := vendingmachine.Solution(5, 0.99)
		assert.Equal(t, []int{1, 0, 0, 0, 0, 4}, res)
	})

	t.Run("Test Solution M=3.14, P=1.99", func(t *testing.T) {
		res := vendingmachine.Solution(3.14, 1.99)
		assert.Equal(t, []int{0, 1, 1, 0, 0, 1}, res)
	})

	t.Run("Test Solution M=3.14, P=1.99", func(t *testing.T) {
		res := vendingmachine.Solution(3, 0.01)
		assert.Equal(t, []int{4, 0, 2, 1, 1, 2}, res)
	})

	t.Run("Test Solution M=3.14, P=1.99", func(t *testing.T) {
		res := vendingmachine.Solution(4, 3.14)
		assert.Equal(t, []int{1, 0, 1, 1, 1, 0}, res)
	})

	t.Run("Test Solution M=3.14, P=1.99", func(t *testing.T) {
		res := vendingmachine.Solution(0.45, 0.34)
		assert.Equal(t, []int{1, 0, 1, 0, 0, 0}, res)
	})

}
