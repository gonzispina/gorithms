package merge2sortedarraysinplace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test equal size arrays", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		merge(nums1, 3, nums2, 3)
		assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
	})

	t.Run("Test nums2 lesser elements ", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 6, 6, 6, 6, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		merge(nums1, 4, nums2, 3)
		assert.Equal(t, []int{1, 2, 2, 3, 5, 6, 6, 6, 6, 6}, nums1)
	})
}
