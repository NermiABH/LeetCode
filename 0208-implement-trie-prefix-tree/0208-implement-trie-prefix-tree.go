type Trie struct {
	nodes [26]*Trie
    hasEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	trie := this
	for i := 0; i < len(word); i++ {
		b := word[i] - 'a'
		if trie.nodes[b] == nil {
			trie.nodes[b] = &Trie{};
		}
        trie = trie.nodes[b]
	}
    trie.hasEnd = true
}

func (this *Trie) Search(word string) bool {
    trie := this
	for i := 0; i < len(word); i++ {
		b := word[i] - 'a'
		if trie.nodes[b] == nil {
			return false
		}
        trie = trie.nodes[b]
	}
    return trie.hasEnd
}

func (this *Trie) StartsWith(prefix string) bool {
    trie := this
	for i := 0; i < len(prefix); i++ {
		b := prefix[i] - 'a'
		if trie.nodes[b] == nil {
			return false
		}
        trie = trie.nodes[b]
	}
    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
