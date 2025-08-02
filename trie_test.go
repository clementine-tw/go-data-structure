package main

import (
	"testing"
)

func TestInsert(t *testing.T) {

	trie := NewTrie()

	keys := []string{"toney", "tone"}
	for _, key := range keys {
		trie.Insert(key)
	}

	for _, key := range keys {
		cur := trie
		for _, r := range key {
			got, ok := cur.Children[r]
			if !ok {
				t.Errorf("Missing rune of word %q in Trie", key)
			}
			cur = got
		}
		if !cur.IsEnd {
			t.Errorf("The last rune of word %q should be set IsEnd to true", key)
		}
	}
}

func TestSearch(t *testing.T) {

	trie := NewTrie()

	keys := []string{"toney", "tone"}
	for _, key := range keys {
		trie.Insert(key)
	}

	t.Run("Should find keys", func(t *testing.T) {
		for _, key := range keys {
			if !trie.Search(key) {
				t.Errorf("%q not found", key)
			}
		}
	})

	t.Run("Shouldn't find key", func(t *testing.T) {
		notExist := "tonee"
		if trie.Search(notExist) {
			t.Errorf("%q shouldn't be found", notExist)
		}
	})
}

func TestIsPrefix(t *testing.T) {

	trie := NewTrie()

	keys := []string{"toney", "tone"}
	for _, key := range keys {
		trie.Insert(key)
	}

	t.Run("Partial prefix", func(t *testing.T) {
		prefix := "ton"
		if !trie.IsPrefix(prefix) {
			t.Errorf("%q should be prefix", prefix)
		}
	})

	t.Run("Full word as prefix", func(t *testing.T) {
		prefix := "tone"
		if !trie.IsPrefix(prefix) {
			t.Errorf("%q should be prefix", prefix)
		}
	})

	t.Run("Not prefix", func(t *testing.T) {
		prefix := "tan"
		if trie.IsPrefix(prefix) {
			t.Errorf("%q shouldn't be prefix", prefix)
		}
	})
}
