package sorting_test

import (
	"github.com/gonzispina/gorithms/sorting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingSort(t *testing.T) {
	t.Run("Test Solution A = 9 3 3 5 6 0 1 4", func(t *testing.T) {
		arr := []int{9, 3, 3, 5, 6, 0, 1, 4}
		res := sorting.CountSort(arr, 1)
		assert.Equal(t, []int{0, 1, 3, 3, 4, 5, 6, 9}, res)
	})

	t.Run("Test Solution A = 9 3 3 5 6 0 1 4 4", func(t *testing.T) {
		arr := []int{9, 3, 3, 5, 6, 0, 1, 4, 4}
		res := sorting.CountSort(arr, 1)
		assert.Equal(t, []int{0, 1, 3, 3, 4, 4, 5, 6, 9}, res)
	})
}
