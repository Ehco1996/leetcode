/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (44.44%)
 * Likes:    416
 * Dislikes: 0
 * Total Accepted:    34K
 * Total Submissions: 74.3K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 * 示例:
 *
*  LRUCache cache = new LRUCache(2);
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 *
 *
*/

// @lc code=start
type Node struct {
	val  int
	prev *Node
	next *Node
	key  int
}

type DBL struct {
	root *Node
	lens int
}

func NewDBL(lens int) DBL {
	node := &Node{-1, nil, nil}
	node.prev = node
	node.next = node

	return DBL{
		root: node,
		lens: lens,
	}
}

func (d *DBL) Head() *Node {
	return d.root.next
}

func (d *DBL) Tail() *Node {
	return d.root.prev
}

func (d *DBL) append(node *Node) {
	d.Tail().next = node
	node.prev = d.Tail()
	d.root.prev = node
	d.lens++
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
	d.lens--
}

func (d *DBL) moveToTail(node *Node) {
	d.remove(node)
	d.append(node)
}

type LRUCache struct {
	capacity int
	dbl      DBL
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity,
		dbl:   NewDBL(0),
		cache: make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) append(k, v int) {
	node := &Node{v, nil, nil, k}
	this.dbl.append(node)
	this.cache[k] = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if ok {
		this.dbl.moveToTail(node)
		return node.val
	}
	return nil
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		this.dbl.moveToTail(node)
	} else {
		if this.capacity == this.dbl.lens {
			head := this.dbl.Head()
			delete(this.cache, head.key)
			this.dbl.remove(head)
		}
		this.append(key, value)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

