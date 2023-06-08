package mediansortedarrays_test

import (
	"github.com/gonzispina/gorithms/mediansortedarrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N=2, M=1", func(t *testing.T) {
		a := []int{1, 3}
		b := []int{2}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(2), res)
	})

	t.Run("Test B bigger than A, m = 1", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{5}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, 2.5, res)
	})

	t.Run("Test B bigger than A, m = 2", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{5, 6}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(3), res)
	})

	t.Run("Test A bigger than B, m = 1", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{-1}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, 1.5, res)
	})

	t.Run("Test A bigger than B, m = 2", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{-2, -1}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(1), res)
	})

	t.Run("B bigger than A, negative numbers", func(t *testing.T) {
		a := []int{-2, -1}
		b := []int{3}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(-1), res)
	})

	t.Run("Test Solution Fully contained", func(t *testing.T) {
		a := []int{2, 2, 4, 4}
		b := []int{2, 2, 4, 4}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(3), res)
	})

	t.Run("Test Solution Fully contained B odds", func(t *testing.T) {
		a := []int{2, 2, 4, 4}
		b := []int{2, 2, 4}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(2), res)
	})

	t.Run("Some overlapping", func(t *testing.T) {
		a := []int{2, 3, 5}
		b := []int{1, 3}

		res := mediansortedarrays.Solution(a, b)
		assert.Equal(t, float64(3), res)
	})
}
