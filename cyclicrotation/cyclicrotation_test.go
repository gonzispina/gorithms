package cyclicrotation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/cyclicrotation"
)

func TestCyclicRotation(test *testing.T) {
	test.Run("TestCyclicRotation - Nil array", func(t *testing.T) {
		res := cyclicrotation.CyclicRotation(nil, 0)
		assert.Equal(t, []int{}, res)
	})

	test.Run("TestCyclicRotation - Empty array", func(t *testing.T) {
		res := cyclicrotation.CyclicRotation([]int{}, 0)
		assert.Equal(t, []int{}, res)
	})

	test.Run("TestCyclicRotation - K = 0", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 0)
		assert.Equal(t, a, res)
	})

	test.Run("TestCyclicRotation - K = 1", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 1)
		assert.Equal(t, []int{6, 1, 2, 3, 4, 5}, res)
	})

	test.Run("TestCyclicRotation - K = 2", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 2)
		assert.Equal(t, []int{5, 6, 1, 2, 3, 4}, res)
	})

	test.Run("TestCyclicRotation - K = 3 = N / 2", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 3)
		assert.Equal(t, []int{4, 5, 6, 1, 2, 3}, res)
	})

	test.Run("TestCyclicRotation - K = 4 > N / 2", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 4)
		assert.Equal(t, []int{3, 4, 5, 6, 1, 2}, res)
	})

	test.Run("TestCyclicRotation - K = 6 = N", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 6)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, res)
	})

	test.Run("TestCyclicRotation - K = 7 > N", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 7)
		assert.Equal(t, []int{6, 1, 2, 3, 4, 5}, res)
	})

	test.Run("TestCyclicRotation - K = 8 > N", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 8)
		assert.Equal(t, []int{5, 6, 1, 2, 3, 4}, res)
	})

	test.Run("TestCyclicRotation - K = 10 > N", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 10)
		assert.Equal(t, []int{3, 4, 5, 6, 1, 2}, res)
	})

	test.Run("TestCyclicRotation - K = 11 > N", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 11)
		assert.Equal(t, []int{2, 3, 4, 5, 6, 1}, res)
	})

	test.Run("TestCyclicRotation - K = 12 > N", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		res := cyclicrotation.CyclicRotation(a, 12)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, res)
	})

	test.Run("TestCyclicRotation - K = 100000000", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		start := time.Now()
		res := cyclicrotation.CyclicRotation(a, 100000000000000)
		diff := time.Now().Sub(start).Seconds()
		assert.Equal(t, []int{3, 4, 5, 6, 1, 2}, res)
		assert.True(t, diff < 0.00001)
	})
}
