package maxcounters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/maxcounters"
)

func TestMaxCounters(test *testing.T) {
	test.Run("TestMaxCounters - N = 0", func(t *testing.T) {
		res := maxcounters.MaxCounters(0, []int{1, 2, 3, 4})
		assert.Equal(t, []int{}, res)
	})

	test.Run("TestMaxCounters - N = 5. A = nil", func(t *testing.T) {
		res := maxcounters.MaxCounters(5, nil)
		assert.Equal(t, []int{0, 0, 0, 0, 0}, res)
	})

	test.Run("TestMaxCounters - N = 5. A = empty", func(t *testing.T) {
		res := maxcounters.MaxCounters(5, []int{})
		assert.Equal(t, []int{0, 0, 0, 0, 0}, res)
	})

	test.Run("TestMaxCounters - N = 1. A = Zero array", func(t *testing.T) {
		a := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		res := maxcounters.MaxCounters(1, a)
		assert.Equal(t, []int{0}, res)
	})

	test.Run("TestMaxCounters - N = 1. A = Zeroes, negatives and n plus 1 numbers", func(t *testing.T) {
		a := []int{0, 0, -1, 0, -4, 0, 10, 0, 4, 0}
		res := maxcounters.MaxCounters(1, a)
		assert.Equal(t, []int{0}, res)
	})

	test.Run("TestMaxCounters - N = 5. One maximum", func(t *testing.T) {
		a := []int{3, 4, 4, 6, 1, 4, 4}
		res := maxcounters.MaxCounters(5, a)
		assert.Equal(t, []int{3, 2, 2, 4, 2}, res)
	})
}
