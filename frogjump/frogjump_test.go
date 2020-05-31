package frogjump_test

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/frogjump"
)


func TestFrogJump(t *testing.T) {
	t.Run("TestFrogJump - X = 0, Y = 0, D = 1", func(test *testing.T) {
		res := frogjump.FrogJump(1, 1, 1)
		assert.Equal(t, 0, res)
	})

	t.Run("TestFrogJump - X = 0, Y = 1, D = 0", func(test *testing.T) {
		res := frogjump.FrogJump(0, 1, 0)
		assert.Equal(t, math.MaxInt64, res)
	})

	t.Run("TestFrogJump - X = 0, Y = 1, D = 1", func(test *testing.T) {
		res := frogjump.FrogJump(0, 1, 1)
		assert.Equal(t, 1, res)
	})

	t.Run("TestFrogJump - X = 0, Y = Max possible int, D = 1", func(test *testing.T) {
		start := time.Now()
		res := frogjump.FrogJump(0, math.MaxInt64, 1)
		diff := time.Now().Sub(start).Seconds()
		assert.Equal(t, math.MaxInt64, res)
		assert.True(t, diff < 0.000001)
	})
}

