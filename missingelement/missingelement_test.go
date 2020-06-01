package missingelement_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/missingelement"
)

func TestMissingElement(test *testing.T) {
	test.Run("TestMissingElement - Nil array", func(t *testing.T) {
		res := missingelement.MissingElement(nil)
		assert.Equal(t, 0, res)
	})

	test.Run("TestMissingElement - Empty array", func(t *testing.T) {
		res := missingelement.MissingElement([]int{})
		assert.Equal(t, 0, res)
	})

	test.Run("TestMissingElement - Array of ten missing seven", func(t *testing.T) {
		res := missingelement.MissingElement([]int{2, 4, 3, 5, 6, 1, 9 ,8})
		assert.Equal(t, 7, res)
	})

	test.Run("TestMissingElement - Array of ten missing missing two", func(t *testing.T) {
		res := missingelement.MissingElement([]int{7, 4, 3, 5, 6, 1, 9 ,8})
		assert.Equal(t, 2, res)
	})

	test.Run("TestMissingElement - Array of five missing missing four", func(t *testing.T) {
		res := missingelement.MissingElement([]int{2, 3, 5, 1})
		assert.Equal(t, 4, res)
	})
}
