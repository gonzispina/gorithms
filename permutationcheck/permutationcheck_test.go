package permutationcheck_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/permutationcheck"
)

func TestPermutationCheck(test *testing.T) {
	test.Run("TestPermutationCheck - Array is nil", func(t *testing.T) {
		res := permutationcheck.PermutationCheck(nil)
		assert.False(t, res)
	})

	test.Run("TestPermutationCheck - Array is empty", func(t *testing.T) {
		res := permutationcheck.PermutationCheck(nil)
		assert.False(t, res)
	})

	test.Run("TestPermutationCheck - N = 1 Permutation", func(t *testing.T) {
		a := []int{1}
		res := permutationcheck.PermutationCheck(a)
		assert.True(t, res)
	})

	test.Run("TestPermutationCheck - N = 1 Not a permutation", func(t *testing.T) {
		a := []int{2}
		res := permutationcheck.PermutationCheck(a)
		assert.False(t, res)
	})

	test.Run("TestPermutationCheck - N = 5 Permutation", func(t *testing.T) {
		a := []int{1, 4, 3, 5, 2}
		res := permutationcheck.PermutationCheck(a)
		assert.True(t, res)
	})

	test.Run("TestPermutationCheck - N = 5 Not Permutation", func(t *testing.T) {
		a := []int{1, 4, 6, 5, 2}
		res := permutationcheck.PermutationCheck(a)
		assert.False(t, res)
	})

	test.Run("TestPermutationCheck - N = 99999999999 Not Permutation", func(t *testing.T) {
		// Takes a lot of time allocating the memory
		a := make([]int, 999999999)

		start := time.Now()
		res := permutationcheck.PermutationCheck(a)
		diff := time.Now().Sub(start).Milliseconds()

		assert.False(t, res)
		assert.True(t, diff < 15)
	})

	test.Run("TestPermutationCheck - N = 3 Sum OK but not a permutation", func(t *testing.T) {
		a := []int{2, 2, 2}
		res := permutationcheck.PermutationCheck(a)
		assert.False(t, res)
	})
}
