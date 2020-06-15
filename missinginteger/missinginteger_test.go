package missinginteger_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/missinginteger"
)

func TestMissingInteger(test *testing.T) {
	test.Run("TestMissingInteger - Nil array", func(t *testing.T) {
		res := missinginteger.MissingInteger(nil)
		assert.Equal(t, 1, res)
	})

	test.Run("TestMissingInteger - Empty array", func(t *testing.T) {
		res := missinginteger.MissingInteger([]int{})
		assert.Equal(t, 1, res)
	})

	test.Run("TestMissingInteger - N = 1 Zero", func(t *testing.T) {
		a := []int{0}
		res := missinginteger.MissingInteger(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestMissingInteger - N = 5 All negatives", func(t *testing.T) {
		a := []int{-1, -2, -3, -4, -5}
		res := missinginteger.MissingInteger(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestMissingInteger - N = 100000 All but 1", func(t *testing.T) {
		var a []int

		for i := 2; len(a) < 100000; i++ {
			a = append(a, i)
		}

		res := missinginteger.MissingInteger(a)
		assert.Equal(t, 1, res)
	})

	test.Run("TestMissingInteger - N = 100000 All from 1 to 100000", func(t *testing.T) {
		var a []int

		for i := 1; len(a) < 100000; i++ {
			a = append(a, i)
		}

		start := time.Now()
		res := missinginteger.MissingInteger(a)
		diff := time.Now().Sub(start).Milliseconds()

		assert.Equal(t, 100001, res)
		assert.True(t, diff < 2)
	})

	test.Run("TestMissingInteger - N = 100000 All from 0 to 100000", func(t *testing.T) {
		var a []int

		for i := 0; len(a) < 100000; i++ {
			a = append(a, i)
		}

		start := time.Now()
		res := missinginteger.MissingInteger(a)
		diff := time.Now().Sub(start).Milliseconds()

		assert.Equal(t, 100000, res)
		assert.True(t, diff < 2)
	})

	test.Run("TestMissingInteger - K within 0 and 99998 union 100001 and 100002", func(t *testing.T) {
		var a []int

		for i := 0; len(a) < 99999; i++ {
			a = append(a, i)
		}

		a = append(a, 100001, 100002)

		start := time.Now()
		res := missinginteger.MissingInteger(a)
		diff := time.Now().Sub(start).Milliseconds()

		assert.Equal(t, 99999, res)
		assert.True(t, diff < 2)
	})
}
