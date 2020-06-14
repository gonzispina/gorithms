package frogriverone_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/frogriverone"
)

func TestFrogRiverOne(test *testing.T) {
	test.Run("TestFrogRiverOne - Nil Array", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(1, nil)
		assert.Equal(t, -1, res)
	})

	test.Run("TestFrogRiverOne - Empty array", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(1, []int{})
		assert.Equal(t, -1, res)
	})

	test.Run("TestFrogRiverOne - X = 0", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(0, []int{})
		assert.Equal(t, 0, res)
	})

	test.Run("TestFrogRiverOne - X < 0", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(-1, []int{})
		assert.Equal(t, -1, res)
	})

	test.Run("TestFrogRiverOne - X > A length", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(2, []int{1})
		assert.Equal(t, -1, res)
	})

	test.Run("TestFrogRiverOne - X = A length", func(t *testing.T) {
		a := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		res := frogriverone.FrogRiverOne(10, a)
		assert.Equal(t, len(a), res)
	})

	test.Run("TestFrogRiverOne - X = 5, Path complete in 6", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4})
		assert.Equal(t, 6, res)
	})

	test.Run("TestFrogRiverOne - X = 10, Path complete in 20", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(10, []int{5, 4, 6, 1, 5, 0, 3, 5, 8, 7, 6, 8, 2, 4, 6, 7, 2, 9, 2, 0, 10, 1, 3, 1, 4, 2, 3})
		assert.Equal(t, 20, res)
	})

	test.Run("TestFrogRiverOne - X = 5, Path is never completed", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 0, 4})
		assert.Equal(t, -1, res)
	})

	test.Run("TestFrogRiverOne - X = 1", func(t *testing.T) {
		res := frogriverone.FrogRiverOne(1, []int{1})
		assert.Equal(t, 0, res)
	})
}
