package sortmovietuples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		movies := []string{
			"Star Wars",
			"El señor de los anillos",
			"Volver al futuro",
		}

		chars := []char{
			{"Leia Organa/Skywalker", "Star wars", 7},
			{"Frodo Baggins", "El señor de los anillos", 2},
			{"Samwise Gamgee (Sam)", "El señor de los anillos", 1},
			{"Luke Skywalker", "Star wars", 4},
			{"Emmett Doc Brown", "Volver al futuro", 1},
			{"Leia Organa/Skywalker", "Star wars", 4},
			{"Samwise Gamgee (Sam)", "El señor de los anillos", 1},
			{"Emmett Doc Brown", "Volver al futuro", 1},
		}

		expected := []char{
			{"Samwise Gamgee (Sam)", "El señor de los anillos", 1},
			{"Frodo Baggins", "El señor de los anillos", 2},
			{"Luke Skywalker", "Star wars", 4},
			{"Leia Organa/Skywalker", "Star wars", 4},
			{"Leia Organa/Skywalker", "Star Wars", 7},
			{"Emmett Doc Brown", "Volver al futuro", 1},
		}

		res := sortMovies(movies, chars)
		assert.Equal(t, expected, res)
	})
}
