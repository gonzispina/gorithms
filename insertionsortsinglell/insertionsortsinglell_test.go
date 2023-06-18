package insertionsortsinglell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N= ", func(t *testing.T) {
		node := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: -1,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		}

		res := insertionSortList(node)
		assert.Equal(t, &ListNode{
			Val: -1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		}, res)
	})
}
