package primecount

import "math/big"

func getPrimorials() []*big.Int {
	primes := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	res := []*big.Int{big.NewInt(6)}
	lastIndex := 0

	for i := 2; i < len(primes); i++ {
		last := res[lastIndex]
		prime := primes[i]
		res = append(res, big.NewInt(0).Mul(last, big.NewInt(prime)))
		lastIndex++
	}

	return res
}

func Solution(N int64) int32 {
	if N == 1 {
		return 0
	}

	primorials := getPrimorials()
	res := int32(0)
	n := big.NewInt(N)

	for index, primorial := range primorials {
		if primorial.Cmp(n) <= 0 {
			continue
		}

		res = int32(index + 1)
		break
	}

	return res
}
