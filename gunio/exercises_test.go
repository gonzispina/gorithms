package gunio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEpigramOnFailure(t *testing.T) {
	t.Run("Test 'why and how'", func(t *testing.T) {
		res := EpigramOnFailure("why and how")
		assert.Equal(t, 569, res)
	})

	t.Run("Test the EpigramOfFailure", func(t *testing.T) {
		text := "Dealing with failure is easy: Work hard to improve. Success is also easy to handle: You’ve solved the wrong problem. Work hard to improve."
		res := EpigramOnFailure(text)
		assert.Equal(t, 2865, res)
	})
}

func TestFibonacciOddSum(t *testing.T) {
	t.Run("Test first 10 (0, 1, 1, 2, 3, 5, 8, 13, 21, 34)", func(t *testing.T) {
		res := FibonacciOddSum(35)
		assert.Equal(t, 44, res)
	})

	t.Run("Test first all less than 10000", func(t *testing.T) {
		res := FibonacciOddSum(10000)
		assert.Equal(t, 14328, res)
	})
}

func TestLegionnaries(t *testing.T) {
	t.Run("Test first 13 digits", func(t *testing.T) {
		// 1 I , 2 II , 3 III, 4 IV, 5 V, 6 VI, 7 VII, 8 VIII, 9 IX, 10 X, 11 XI, 12 XII, 13 XIII
		res := Legionnaries(13)
		assert.Equal(t, 5, res)
	})

	t.Run("Test first 2660 digits", func(t *testing.T) {
		res := Legionnaries(2660)
		assert.Equal(t, 3977, res)
	})
}
