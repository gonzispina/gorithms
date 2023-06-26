package sortvaluesbyranking

import "github.com/gonzispina/gorithms/trie"

type indexCount struct {
	index int
	count int
}

func insertIndexCount(dicc *trie.Trie, key string, index int) {
	dicc.Insert(key, &indexCount{index: index, count: 0})
}

func updateIndexCount(dicc *trie.Trie, key string, count int) {
	i, _ := dicc.Search(key)
	ic := i.(*indexCount)
	ic.count = count
	dicc.Insert(key, ic)
}

func getIndexCount(dicc *trie.Trie, key string) (int, int) {
	i, _ := dicc.Search(key)
	ic := i.(*indexCount)
	return ic.index, ic.count
}

func sort(ranking []string, values []string) []string {
	dicc := trie.NewTrie()

	for i, r := range ranking { // O(len(ranking))
		insertIndexCount(dicc, r, i)
	}

	for _, v := range values { // O(len(values))
		_, c := getIndexCount(dicc, v) // O(1)
		updateIndexCount(dicc, v, c+1)
	}

	indexes := make([]int, len(ranking))
	for _, r := range ranking {
		i, count := getIndexCount(dicc, r)

		if i == 0 {
			indexes[i] = count - 1
			continue
		}

		indexes[i] = count + indexes[i-1]
	}

	res := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		index, _ := getIndexCount(dicc, values[i])
		res[indexes[index]] = ranking[index]
		indexes[index]--
	}

	return res
}
