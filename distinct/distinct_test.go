package distinct_test

import (
	"github.com/gonzispina/gorithms/distinct"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestDistinct(t *testing.T) {
	t.Run("Nil array", func(t *testing.T) {
		v := distinct.Distinct(nil)
		assert.Equal(t, 0, v)
	})

	t.Run("All equal numbers in the array", func(t *testing.T) {
		v := distinct.Distinct([]int{0, 0, 0, 0, 0})
		assert.Equal(t, 1, v)
	})

	t.Run("All different numbers", func(t *testing.T) {
		v := distinct.Distinct([]int{1, 2, 3, 4, 0})
		assert.Equal(t, 5, v)
	})

	t.Run("Negative numbers", func(t *testing.T) {
		v := distinct.Distinct([]int{-1, 2, 1, 3, 4, 0})
		assert.Equal(t, 6, v)
	})

	t.Run("Some random numbers", func(t *testing.T) {
		v := distinct.Distinct([]int{5, 10, 10, 5, 4, 3, 2, 0, 2, 9})
		assert.Equal(t, 7, v)
	})

}
