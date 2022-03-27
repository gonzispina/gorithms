package gunio

func isLetter(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func isVowel(s rune) bool {
	return (s == 'a') ||
		(s == 'e') ||
		(s == 'i') ||
		(s == 'o') ||
		(s == 'u') ||
		(s == 'A') ||
		(s == 'E') ||
		(s == 'I') ||
		(s == 'O') ||
		(s == 'U')
}

func EpigramOnFailure(text string) int {
	var sum, factor int
	for _, s := range text {
		if !isLetter(s) {
			continue
		}

		factor = 1
		if isVowel(s) {
			factor = -1
		}

		sum += factor * int(s)
	}

	return sum
}

func FibonacciOddSum(N int) int {
	previous := make(map[int]int)
	previous[1] = 1
	previous[2] = 1

	sum := 2
	for i := 3; previous[i-1]+previous[i-2] <= N; i++ {
		fib := previous[i-1] + previous[i-2]
		previous[i] = fib
		if fib%2 != 0 {
			sum += fib
		}
	}

	return sum
}

func Legionnaries(N int) int {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	xCountValues := []int{
		0, // "M"
		0, // "CM"
		0, // "D"
		0, // "CD"
		0, // "C"
		1, // "XC"
		0, // "L"
		1, // "XL"
		1, // "X",
		1, // "IX"
		0, // "V"
		0, // "IV"
		0, // "I"
	}

	sum := 0
	for i := 1; i <= N; i++ {
		num := i
		for j := 0; j < len(values); j++ {
			for num >= values[j] {
				num -= values[j]
				sum += xCountValues[j]
			}
		}
	}

	return sum
}
