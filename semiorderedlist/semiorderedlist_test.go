package semiorderedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		arr := []int{3, 17, 2, 20, 1, 23, 5, 11, 9}
		sort(arr, 4, 12)
		assert.Equal(t, []int{1, 2, 3, 5, 9, 11, 17, 20, 23}, arr)
	})
}
