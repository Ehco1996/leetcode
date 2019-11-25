/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 *
 * https://leetcode-cn.com/problems/design-hashset/description/
 *
 * algorithms
 * Easy (54.47%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    7.6K
 * Total Submissions: 13.9K
 * Testcase Example:  '["MyHashSet","add","add","contains","contains","add","contains","remove","contains"]\n' +
  '[[],[1],[2],[1],[3],[2],[2],[2],[2]]'
 *
 * 不使用任何内建的哈希表库设计一个哈希集合
 *
 * 具体地说，你的设计应该包含以下的功能
 *
 *
 * add(value)：向哈希集合中插入一个值。
 * contains(value) ：返回哈希集合中是否存在这个值。
 * remove(value)：将给定值从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
 *
 *
 *
 * 示例:
 *
 * MyHashSet hashSet = new MyHashSet();
 * hashSet.add(1);
 * hashSet.add(2);
 * hashSet.contains(1);    // 返回 true
 * hashSet.contains(3);    // 返回 false (未找到)
 * hashSet.add(2);
 * hashSet.contains(2);    // 返回 true
 * hashSet.remove(2);
 * hashSet.contains(2);    // 返回  false (已经被删除)
 *
 *
 *
 * 注意：
 *
 *
 * 所有的值都在 [1, 1000000]的范围内。
 * 操作的总数目在[1, 10000]范围内。
 * 不要使用内建的哈希集合库。
 *
 *
*/

// @lc code=start
type MyHashSet struct {
	Store []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	store := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		store[i] = -1
	}

	return MyHashSet{Store: store}
}

func (this *MyHashSet) Add(key int) {
	hKey := key % 100000
	this.Store[hKey] = key
}

func (this *MyHashSet) Remove(key int) {
	hKey := key % 100000
	this.Store[hKey] = -1
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	hKey := key % 100000
	return this.Store[hKey] != -1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end

