package main

import (
	"fmt"

	"github.com/rugggger/go-search/internal/trie"
)

func main() {
	dictionary := trie.NewTrie()
	dictionary.Insert("ap")
	dictionary.Insert("apple")
	println("---")

	fmt.Println(dictionary.Exists("ap"))
	fmt.Println(dictionary.Exists("apple"))

}
