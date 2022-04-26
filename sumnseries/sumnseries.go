package sumnseries

func Solution(N int64) int32 {
	m := int64(1000000007)
	return int32(((N % m) * (N % m)) % m)
}
