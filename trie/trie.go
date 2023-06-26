package trie

type Node struct {
	defined    bool
	nodes      *[26]*Node
	definition interface{}
}

func newNode() *Node {
	return &Node{
		defined: false,
		nodes:   nil,
	}
}

func NewTrie() *Trie {
	return &Trie{root: newNode()}
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(key string, value interface{}) {
	n := t.root
	for len(key) > 0 {
		char := key[0]
		key = key[1:]
		index := int(char) % 26

		if n.nodes == nil {
			n.nodes = &[26]*Node{}
		}

		if n.nodes[index] == nil {
			n.nodes[index] = newNode()
		}

		n = n.nodes[index]
	}
	n.defined = true
	n.definition = value
}

func (t *Trie) Search(word string) (interface{}, bool) {
	n := t.root
	for len(word) > 0 {
		char := word[0]
		word = word[1:]
		index := int(char) % 26

		if n.nodes == nil || n.nodes[index] == nil {
			return nil, false
		}

		n = n.nodes[index]
	}

	return n.definition, n.defined
}
