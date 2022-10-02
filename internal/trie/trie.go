package trie

import "fmt"

type trie_node struct {
	children [26]*trie_node
	word     string
	wordEnds bool
}
type Trie struct {
	root *trie_node
}

func NewTrie() *Trie {
	t := new(Trie)
	t.root = new(trie_node)
	return t
}

func (t *Trie) Insert(word string) {

	fmt.Println("insert", word)
	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.children[index] == nil {
			current.children[index] = new(trie_node)
		}
		current = current.children[index]
	}
	current.word = word
	current.wordEnds = true
}

func (t *Trie) Exists(word string) bool {
	fmt.Printf("searching %s\n", word)

	current := t.root
	for _, wr := range word {
		index := wr - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current.wordEnds

}

func Search() {
	fmt.Print("search")
}
