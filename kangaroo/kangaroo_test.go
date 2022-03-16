package kangaroo_test

import (
	"github.com/gonzispina/gorithms/kangaroo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKangaroo(t *testing.T) {
	/*
	   Escenario 1: el canguro uno está por delante del dos (x1 - x2 > 0)
	       Posible si el canguro uno es más rápido que el canguro dos (v2 - v1 < 0) => n < 0
	       Imposible si el canguro dos es más rapido que el uno (v2 - v1 > 0) => n > 0
	*/
	t.Run("Escenario 1 Posible", func(t *testing.T) {
		res := kangaroo.Kangaroo(0, 3, 4, 2)
		assert.Equal(t, "YES", res)
	})

	t.Run("Escenario 1 Imposible", func(t *testing.T) {
		res := kangaroo.Kangaroo(4, 3, 0, 2)
		assert.Equal(t, "NO", res)
	})

	/*
	   Escenario 2: el canguro dos está por delante del uno (x1 - x2 < 0)
	       Posible si el canguro uno es más rapido que el canguro dos (v2 - v1 < 0) => n > 0
	       Imposible si el canguro dos es más rápido que el canguro uno (v2 - v1 > 0) => n < 0
	*/
	t.Run("Escenario 2 Posible", func(t *testing.T) {
		res := kangaroo.Kangaroo(4, 2, 0, 3)
		assert.Equal(t, "YES", res)
	})

	t.Run("Escenario 2 Imposible", func(t *testing.T) {
		res := kangaroo.Kangaroo(21, 6, 47, 3)
		assert.Equal(t, "NO", res)
	})
}
