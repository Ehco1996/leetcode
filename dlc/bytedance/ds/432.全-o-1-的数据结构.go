/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 *
 * https://leetcode-cn.com/problems/all-oone-data-structure/description/
 *
 * algorithms
 * Hard (35.16%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    2.2K
 * Total Submissions: 6.4K
 * Testcase Example:  '["AllOne","getMaxKey","getMinKey"]\n[[],[],[]]'
 *
 * 实现一个数据结构支持以下操作：
 *
 *
 * Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
 * Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key
 * 不存在，这个函数不做任何事情。key 保证不为空字符串。
 * GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
 * GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
 *
 *
 * 挑战：以 O(1) 的时间复杂度实现所有操作。
 *
 */

// @lc code=start
type Node struct {
	val  int
	prev *Node
	next *Node
	keys map[string]bool
}

type DBL struct {
	root *Node
}

func NewDBL() DBL {
	node := &Node{-1, nil, nil}
	node.prev = node
	node.next = node

	return DBL{
		root: node,
	}
}

func (d *DBL) Head() *Node {
	return d.root.next
}

func (d *DBL) Tail() *Node {
	return d.root.prev
}

func (d *DBL) remove(node *Node) {
	if node == d.root {
		return
	}

	if node == d.Tail() {
		d.root.prev = node.prev
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev.next = node.next
}

type AllOne struct {
	keyMap map[string]*Node
	db     DBL
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		keyMap: make(map[string]*Node),
		db:     NewDBL(),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	node, ok := this.db.KeyMap[key]
	if ok {
		node.val++
		next := node.next
		this.db.remove(node)
	} else {
		node := &Node{0, nil, nil, make(map[string]bool)}
		this.db.KeyMap[key] = node
	}
	// 找到第一个当前值小的node
	for node.val <= next.val {
		next = next.next
	}
	if next.val == node.val {
		//合并keys
		for k, _ := range node.keys {
			next.keys[k] = true
		}
	} else {
		// 添加新的节点
		this.db.insert(next, node)
	}

}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {

	node, ok := this.db.KeyMap[key]
	if ok {
		node.val--
		prev := node.prev
		this.db.remove(node)
	} else {
		node := &Node{0, nil, nil, make(map[string]bool)}
		this.db.KeyMap[key] = node
	}
	// 找到第一个当前值小的node
	for node.val <= prev.val {
		prev = prev.prev
	}
	if prev.val == node.val {
		//合并keys
		for k, _ := range node.keys {
			next.keys[k] = true
		}
	} else {
		// 添加新的节点
		this.db.insert(prev, node)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	for key, _ := range this.db.Tail().keys {
		return key
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	for key, _ := range this.db.Head().keys {
		return key
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end

