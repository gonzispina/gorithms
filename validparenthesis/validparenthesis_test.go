package validparenthesis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution ()", func(t *testing.T) {
		res := isValid("()")
		assert.Equal(t, true, res)
	})

	t.Run("Test Solution )(", func(t *testing.T) {
		res := isValid(")(")
		assert.Equal(t, false, res)
	})

	t.Run("Test Solution ({})[", func(t *testing.T) {
		res := isValid("({})[")
		assert.Equal(t, false, res)
	})

	t.Run("Test Solution {{{{", func(t *testing.T) {
		res := isValid("{{{{")
		assert.Equal(t, false, res)
	})

	t.Run("Test Solution {{{{}}}]", func(t *testing.T) {
		res := isValid("{{{{}}}]")
		assert.Equal(t, false, res)
	})

	t.Run("Test Solution {{{{()}()}}}", func(t *testing.T) {
		res := isValid("{{{{()}()}}}")
		assert.Equal(t, true, res)
	})
}
