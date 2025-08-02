package main

type TrieNode struct {
	IsEnd    bool
	Children map[rune]*TrieNode
}

// Create new Trie.
func NewTrie() *TrieNode {
	return &TrieNode{
		IsEnd:    false,
		Children: make(map[rune]*TrieNode),
	}
}

// Insert key into Trie.
func (t *TrieNode) Insert(key string) {
	cur := t
	for _, r := range key {
		if child, ok := cur.Children[r]; ok {
			cur = child
			continue
		}
		cur.Children[r] = &TrieNode{
			Children: make(map[rune]*TrieNode),
		}
		cur = cur.Children[r]
	}
	cur.IsEnd = true
}

// Search for if key is in Trie.
func (t *TrieNode) Search(key string) bool {
	cur := t
	for _, r := range key {
		if child, ok := cur.Children[r]; ok {
			cur = child
			continue
		}
		return false
	}
	return cur.IsEnd
}

// Check if the prefix is partial or full word in Trie.
func (t *TrieNode) IsPrefix(prefix string) bool {
	cur := t
	for _, r := range prefix {
		if child, ok := cur.Children[r]; ok {
			cur = child
			continue
		}
		return false
	}
	return true
}
