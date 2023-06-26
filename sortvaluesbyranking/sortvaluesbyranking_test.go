package sortvaluesbyranking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		ranking := []string{"brasil", "argentina", "alemania", "chile", "colombia", "francia"}
		values := []string{
			"argentina", "brasil", "colombia", "francia", "colombia",
			"argentina", "chile", "alemania", "brasil", "chile",
		}

		expected := []string{
			"brasil", "brasil", "argentina", "argentina", "alemania",
			"chile", "chile", "colombia", "colombia", "francia",
		}

		res := sort(ranking, values)
		assert.Equal(t, expected, res)
	})
} 
