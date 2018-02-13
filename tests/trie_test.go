package tests

import (
	"testing"

	ds "../dstructs"
)

func TestTrieInsertAndContains(t *testing.T) {
	var trie ds.Trie
	trie = ds.NewTrie()
	trie.Insert("hello")
	trie.Insert("hell")
	trie.Insert("venom")
	if !trie.Contains("hello") {
		t.Errorf("The Trie does not contain inserted word : hello")
	}
	if !trie.Contains("hell") {
		t.Errorf("The Trie does not contain inserted word : hell")
	}
	if !trie.Contains("venom") {
		t.Errorf("The Trie does not contain inserted word : venom")
	}
	if trie.Contains("salt") {
		t.Errorf("The Trie contains non-inserted word : salt")
	}
	if trie.Contains(" ") {
		t.Errorf("The Trie contains non-inserted word : <empty_space>")
	}
}
