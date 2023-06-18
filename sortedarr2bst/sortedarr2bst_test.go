package sortedarr2bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test simple triplet", func(t *testing.T) {
		var res []int
		sortedArrayToBST([]int{3, 5, 8}).PreorderTree2Arr(&res)
		assert.Equal(t, []int{5, 3, 8}, res)
	})
}
