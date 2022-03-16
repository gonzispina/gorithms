package yearoftiger_test

import (
	"github.com/gonzispina/gorithms/yearoftiger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYearOfTiger(t *testing.T) {
	t.Run("TestCase 1", func(t *testing.T) {
		arr := []string{"aab", "cab", "baa", "baa"}
		v := yearoftiger.YearOfTiger(arr)
		assert.Equal(t, 3, v)
	})

	t.Run("TestCase 2", func(t *testing.T) {
		arr := []string{"zzz", "zbz", "zbz", "dgf"}
		v := yearoftiger.YearOfTiger(arr)
		assert.Equal(t, 2, v)
	})

	t.Run("TestCase 3", func(t *testing.T) {
		arr := []string{"abc", "cba", "cab", "bac", "bca"}
		v := yearoftiger.YearOfTiger(arr)
		assert.Equal(t, 3, v)
	})
}
