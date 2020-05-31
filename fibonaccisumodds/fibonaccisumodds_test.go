package fibonaccisumodds_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/fibonaccisumodds"
)

func TestFibonacciSumOdds(test *testing.T) {
	test.Run("TestFibonacciSumOdds - N = 0", func(t *testing.T) {
			res := fibonaccisumodds.FibonacciSumOdds(0)
			assert.Equal(t, 0, res)
	})

	test.Run("TestFibonacciSumOdds - N = 1", func(t *testing.T) {
		res := fibonaccisumodds.FibonacciSumOdds(1)
		assert.Equal(t, 1, res)
	})

	test.Run("TestFibonacciSumOdds - N = 2", func(t *testing.T) {
		res := fibonaccisumodds.FibonacciSumOdds(2)
		assert.Equal(t, 2, res)
	})

	test.Run("TestFibonacciSumOdds - N = 1597", func(t *testing.T) {
		res := fibonaccisumodds.FibonacciSumOdds(1597)
		assert.Equal(t, 3382, res)
	})

	test.Run("TestFibonacciSumOdds - N = 1600", func(t *testing.T) {
		res := fibonaccisumodds.FibonacciSumOdds(1600)
		assert.Equal(t, 3382, res)
	})

	test.Run("TestFibonacciSumOdds - N = 10000", func(t *testing.T) {
		res := fibonaccisumodds.FibonacciSumOdds(10000)
		assert.Equal(t, 14328, res)
	})
}

