package sortmovietuples

import "sort"

type char struct {
	character string
	movie     string
	episode   int
}

func sortMovies(movies []string, chars []char) []char {
	moviesDicc := map[string][]map[string]bool{}

	// O(G * log(G))
	sort.Strings(movies)

	// O(N)
	for _, c := range chars { // O(N)
		if _, ok := moviesDicc[c.movie]; !ok { // O(1)
			moviesDicc[c.movie] = make([]map[string]bool, 7) // O(1)
		}

		if moviesDicc[c.movie][c.episode-1] == nil {
			moviesDicc[c.movie][c.episode-1] = map[string]bool{}
		}

		moviesDicc[c.movie][c.episode-1][c.character] = true // O(1)
	}

	// O(G*p)
	var res []char
	for _, movie := range movies { // O(G)
		for i := 0; i < 7; i++ {
			episode := moviesDicc[movie][i] // O(1)
			if episode == nil {
				continue
			}

			var characters []string
			for k := range episode { // O(50)
				characters = append(characters, k)
			}

			sort.Strings(characters) // O(50*log(50))

			for _, character := range characters { // O(50)
				res = append(res, char{
					character: character,
					movie:     movie,
					episode:   i + 1,
				})
			}
		}
	}

	return res
}
