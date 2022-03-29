package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	t.Run("Get first N=10", func(t *testing.T) {
		res := Fibonacci(10)
		assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, res)
	})
}
