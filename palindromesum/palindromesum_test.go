package palindromesum_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/palindromesum"
)

func TestPalindromeSum(test *testing.T) {
	test.Run("TestPalindromeSum - N = 0", func(t *testing.T) {
		res := palindromesum.PalindromeSum(0)
		assert.Equal(t, 0, res)
	})

	test.Run("TestPalindromeSum - N = 10", func(t *testing.T) {
		res := palindromesum.PalindromeSum(10)
		assert.Equal(t, 45, res)
	})

	test.Run("TestPalindromeSum - N = 100", func(t *testing.T) {
		res := palindromesum.PalindromeSum(100)
		assert.Equal(t, 540, res)
	})

	test.Run("TestPalindromeSum - N = 1000", func(t *testing.T) {
		res := palindromesum.PalindromeSum(1000)
		assert.Equal(t, 50040, res)
	})

	test.Run("TestPalindromeSum - N = 10000", func(t *testing.T) {
		res := palindromesum.PalindromeSum(10000)
		assert.Equal(t, 545040, res)
	})
}