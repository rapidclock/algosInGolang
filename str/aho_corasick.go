package str

type TrieNode struct {
	symb     string
	isEnd    bool
	children map[string]*TrieNode
	words    []string
	failNode *TrieNode
}

func newTrie() *TrieNode {
	newRoot := new(TrieNode)
	newRoot.children = make(map[string]*TrieNode)
	newRoot.failNode = nil
	newRoot.words = make([]string, 0)
	newRoot.isEnd = false
	return newRoot
}

func (root *TrieNode) insert(word string) {
	if len(word) <= 0 {
		return
	}
	curNode := root
	for _, c := range word {
		letter := string(c)
		_, ok := curNode.children[letter]
		if !ok {
			newNode := newTrie()
			newNode.symb = letter
			curNode.children[letter] = newNode
		}
		curNode = curNode.children[letter]
	}
	curNode.words = append(curNode.words, word)
	curNode.isEnd = true
}

func (root *TrieNode) populateFailNodes() {
	q := make([]*TrieNode, 0)
	for _, c := range root.children {
		c.failNode = root
		q = append(q, c)
	}
	for len(q) != 0 {
		var node *TrieNode
		node, q = q[0], q[1:]
		for _, child := range node.children {
			fNode := node.failNode
			for fNode != nil {
				posNode, ok := fNode.children[child.symb]
				if ok {
					child.failNode = posNode
					child.words = append(child.words, posNode.words...)
					break
				} else {
					fNode = fNode.failNode
				}
			}
			if fNode == nil {
				child.failNode = root
			}
			q = append(q, child)
		}
	}
}

func (root *TrieNode) contains(word string) bool {
	curNode := root
	for _, c := range word {
		letter := string(c)
		_, ok := curNode.children[letter]
		if !ok {
			return false
		}
		curNode = curNode.children[letter]
	}
	return curNode.isEnd
}

func (root *TrieNode) checkOccurrences(sentence string) map[string][]int {
	soln := make(map[string][]int)
	node := root
	for i := 0; i < len(sentence); {
		letter := string(sentence[i])
		if child, ok := node.children[letter]; ok {
			node = child
			for _, word := range node.words {
				soln[word] = append(soln[word], i-len(word)+1)
			}
		} else if node.failNode != nil {
			node = node.failNode
			continue
		}
		i += 1
	}
	return soln
}

func AhoCorasick(sentence string, dict []string) map[string][]int {
	trie := newTrie()
	for _, word := range dict {
		trie.insert(word)
	}
	trie.populateFailNodes()
	return trie.checkOccurrences(sentence)
}
