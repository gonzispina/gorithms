package yearoftiger

func YearOfTiger(T []string) int {
	counter := map[string]int{}
	currentMax := 0

	for _, t := range T {
		perms := map[string]struct{}{}
		perms[t] = struct{}{}
		perms[string(t[1])+string(t[0])+string(t[2])] = struct{}{}
		perms[string(t[0])+string(t[2])+string(t[1])] = struct{}{}

		for perm, _ := range perms {
			if _, ok := counter[perm]; !ok {
				counter[perm] = 0
			}

			counter[perm]++
			if counter[perm] > currentMax {
				currentMax = counter[perm]
			}
		}
	}

	return currentMax
}
