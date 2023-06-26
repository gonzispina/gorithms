package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	t.Run("Test Solution A = 3, 9, 30, 34, 5", func(t *testing.T) {
		arr := []int{3, 9, 30, 34, 5}
		res := largestNumber(arr)
		assert.Equal(t, "9534330", res)
	})

	t.Run("Test Solution A = 3, 95, 30, 34", func(t *testing.T) {
		arr := []int{3, 95, 30, 34}
		res := largestNumber(arr)
		assert.Equal(t, "9534330", res)
	})

	t.Run("Test Solution A = 10, 2", func(t *testing.T) {
		arr := []int{10, 2}
		res := largestNumber(arr)
		assert.Equal(t, "210", res)
	})
}
