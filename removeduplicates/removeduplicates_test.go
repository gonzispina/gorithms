package removeduplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		count := removeDuplicates(arr)
		assert.Equal(t, []int{0, 1, 2, 3, 4}, arr[0:count])
	})
}
