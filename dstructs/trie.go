package dstructs

import (
	"fmt"
)

type TrieNode struct {
	isEndOfWord bool
	children    map[string]*TrieNode
}

type Trie interface {
	Insert(word string)
	Contains(word string) bool
}

func NewTrie() *TrieNode {
	newRoot := new(TrieNode)
	newRoot.children = make(map[string]*TrieNode)
	return newRoot
}

func (node *TrieNode) String() string {
	return fmt.Sprint(node.children, node.isEndOfWord)
}

func (root *TrieNode) Insert(word string) {
	curNode := root
	for _, c := range word {
		letter := string(c)
		_, ok := curNode.children[letter]
		if !ok {
			newChild := NewTrie()
			curNode.children[letter] = newChild
		}
		curNode = curNode.children[letter]
	}
	curNode.isEndOfWord = true
}

func (root *TrieNode) Contains(word string) bool {
	curNode := root
	for _, c := range word {
		letter := string(c)
		_, ok := curNode.children[letter]
		if !ok {
			return false
		}
		curNode = curNode.children[letter]
	}
	return curNode.isEndOfWord
}
