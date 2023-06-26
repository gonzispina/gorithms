package generateparenthesis

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N=2", func(t *testing.T) {
		res := generateParenthesis(2)
		expected := []string{"()()", "(())"}
		
		sort.Strings(expected)
		sort.Strings(res)

		assert.Equal(t, expected, res)
	})

	t.Run("Test Solution N=3", func(t *testing.T) {
		res := generateParenthesis(3)
		expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}

		sort.Strings(expected)
		sort.Strings(res)

		assert.Equal(t, expected, res)
	})

	t.Run("Test Solution N=4", func(t *testing.T) {
		res := generateParenthesis(4)
		expected := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}

		sort.Strings(expected)
		sort.Strings(res)

		assert.Equal(t, expected, res)
	})
}
