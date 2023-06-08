package atoistring2int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution V=\"42\"", func(t *testing.T) {
		res := myAtoi("+42")
		assert.Equal(t, 42, res)
	})

	t.Run("Test Solution V=\"   -42\"", func(t *testing.T) {
		res := myAtoi("   -42")
		assert.Equal(t, -42, res)
	})

	t.Run("Test Solution V=\"4193 with words\"", func(t *testing.T) {
		res := myAtoi("4193 with words")
		assert.Equal(t, 4193, res)
	})

	t.Run("Test Solution V=\"words and 987\"", func(t *testing.T) {
		res := myAtoi("words and 987")
		assert.Equal(t, 0, res)
	})

	t.Run("Test Solution V=\"\"", func(t *testing.T) {
		res := myAtoi("")
		assert.Equal(t, 0, res)
	})

	t.Run("Test Solution V=\"2147483648\"", func(t *testing.T) {
		res := myAtoi("2147483648")
		assert.Equal(t, 2147483647, res)
	})

	t.Run("Test Solution V=\"2147483646\"", func(t *testing.T) {
		res := myAtoi("2147483646")
		assert.Equal(t, 2147483646, res)
	})

	t.Run("Test Solution V=\"-2147483648\"", func(t *testing.T) {
		res := myAtoi("-2147483648")
		assert.Equal(t, -2147483648, res)
	})

	t.Run("Test Solution V=\"2147483648\"", func(t *testing.T) {
		res := myAtoi("2147483648")
		assert.Equal(t, 2147483647, res)
	})
}
