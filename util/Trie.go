package util

type Node struct {
	Char     string
	Children [26]*Node
	IsWord   bool
	Word     string
}

type Trie struct {
	RootNode *Node
}

func NewNode(char string) *Node {
	node := &Node{Char: char, IsWord: false}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}
	return node
}

func NewTrie() *Trie {
	root := NewNode("\000")
	return &Trie{RootNode: root}
}

func (t *Trie) Insert(word string) error {
	current := t.RootNode
	for i := 0; i < len(word); i++ {
		// 'a' == 97, 'b' == 98, etc.
		index := word[i] - 'a'
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(word[i]))
		}
		current = current.Children[index]
	}
	current.IsWord = true
	current.Word = word
	return nil
}

func (t *Trie) SearchWord(word string) bool {
	current := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if current == nil || current.Children[index] == nil {
			return false
		}
	}
	return true
}

func (n *Node) Traverse(char rune) *Node {
	index := char - 'a'
	return n.Children[index]
}
