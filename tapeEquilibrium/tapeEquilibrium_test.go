package tapeEquilibrium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/tapeEquilibrium"
)

func TestTapeEquilibrium(test *testing.T) {
	test.Run("TestTapeEquilibrium - Nil array", func(t *testing.T) {
		res := tapeEquilibrium.TapeEquilibrium(nil)
		assert.Equal(t, 0, res)
	})

	test.Run("TestTapeEquilibrium - Empty array", func(t *testing.T) {
		res := tapeEquilibrium.TapeEquilibrium([]int{})
		assert.Equal(t, 0, res)
	})

	test.Run("TestTapeEquilibrium - N = 1", func(t *testing.T) {
		a := []int{1}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestTapeEquilibrium - N = 2", func(t *testing.T) {
		a := []int{2000, 0}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 2000, res)
	})

	test.Run("TestTapeEquilibrium - Small array all positives", func(t *testing.T) {
		a := []int{3, 1, 2, 4, 3}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestTapeEquilibrium - Small array all negatives", func(t *testing.T) {
		a := []int{-3, -1, -2, -4, -3}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestTapeEquilibrium - Small array some negatives some positives", func(t *testing.T) {
		a := []int{-3, -1, 2, 4, -3}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestTapeEquilibrium - 5050", func(t *testing.T) {
		a := []int{20, -20, 10, 90, 80, 30, -50, 40}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 0, res)
	})

	test.Run("TestTapeEquilibrium - Tail of zeroes", func(t *testing.T) {
		a := []int{2000, 0, 0, 0, 0}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 2000, res)
	})

	test.Run("TestTapeEquilibrium - Head of zeroes", func(t *testing.T) {
		a := []int{0, 0, 0, 0, 2000, 1, 0, 0, 1}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 1998, res)
	})

	test.Run("TestTapeEquilibrium - Head of zeroes", func(t *testing.T) {
		a := []int{-1000, 1000}
		res := tapeEquilibrium.TapeEquilibrium(a)
		assert.Equal(t, 0, res)
	})
}