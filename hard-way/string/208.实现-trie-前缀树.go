/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 *
 * https://leetcode-cn.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (60.05%)
 * Likes:    133
 * Dislikes: 0
 * Total Accepted:    13.7K
 * Total Submissions: 22.2K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
 *
 * 示例:
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // 返回 true
 * trie.search("app");     // 返回 false
 * trie.startsWith("app"); // 返回 true
 * trie.insert("app");
 * trie.search("app");     // 返回 true
 *
 * 说明:
 *
 *
 * 你可以假设所有的输入都是由小写字母 a-z 构成的。
 * 保证所有输入均为非空字符串。
 *
 *
*/

// @lc code=start
type Node struct {
	Hash   map[string]*Node
	IsWord bool
}

func NewNode(isWord bool) *Node {
	return &Node{Hash: make(map[string]*Node), IsWord: isWord}
}

// hash 每个node上有个hash存着到下一层的字母，并且有一个标志位来表示当前路径是否是一个单词
type Trie struct {
	Head *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Head: NewNode(false)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.Head
	for _, s := range word {
		c := string(s)
		if _, ok := cur.Hash[c]; !ok {
			cur.Hash[c] = NewNode(false)
		}
		cur = cur.Hash[c]
	}
	cur.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.Head
	for _, s := range word {
		c := string(s)
		if _, ok := cur.Hash[c]; !ok {
			return false
		}
		cur = cur.Hash[c]
	}
	return cur.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.Head
	for _, s := range prefix {
		c := string(s)
		if _, ok := cur.Hash[c]; !ok {
			return false
		}
		cur = cur.Hash[c]
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
// @lc code=end

