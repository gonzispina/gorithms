package binarygap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gonzispina/gorithms/binarygap"
)

func TestBinaryGap(test *testing.T) {
	test.Run("BinaryGap - V = 1041 = 10000010001", func(t *testing.T) {
		res := binarygap.BinaryGap(1041)
		assert.Equal(t, 5, res)
	})

	test.Run("BinaryGap - V = 210002 = 110011010001010010", func(t *testing.T) {
		res := binarygap.BinaryGap(210002)
		assert.Equal(t, 3, res)
	})

	test.Run("BinaryGap - V = 393248 = 1100000000000100000", func(t *testing.T) {
		res := binarygap.BinaryGap(393248)
		assert.Equal(t, 11, res)
	})
}
