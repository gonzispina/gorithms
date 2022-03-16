package kangaroo

import "math"

/*
   Escenario 1: el canguro uno está por delante del dos (x1 - x2 > 0)
       Posible si el canguro uno es más rápido que el canguro dos (v2 - v1 < 0) => n < 0
       Imposible si el canguro dos es más rapido que el uno (v2 - v1 > 0) => n > 0

   Escenario 2: el canguro dos está por delante del uno (x1 - x2 < 0)
       Posible si el canguro uno es más rapido que el canguro dos (v2 - v1 < 0) => n > 0
       Imposible si el canguro dos es más rápido que el canguro uno (v2 - v1 > 0) => n < 0
*/

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 == x2 {
		if v2 == v1 {
			return "YES"
		}
		return "NO"
	}

	if v2 == v1 {
		return "NO"
	}

	n := (float64(x1) - float64(x2)) / (float64(v2) - float64(v1))
	if n > 0 && n == math.Floor(n) {
		return "YES"
	}

	return "NO"
}
