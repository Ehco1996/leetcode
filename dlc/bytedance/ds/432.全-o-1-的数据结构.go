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
// TODO
type AllOne struct {
	Store  map[string]int
	Max    int
	MaxKey string
	Min    int
	MinKey string
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{Store: make(map[string]int)}
}

func (this *AllOne) balance(key string) {
	if len(this.Store) == 1 {
		this.Max = this.Store[key]
		this.MaxKey = key
		this.Min = this.Store[key]
		this.MinKey = key
	}
	val := this.Store[key]

	if val < this.Min {
		this.Min = this.Store[key]
		this.MinKey = key

	}
	if va > this.Max {
		this.Max = this.Store[key]
		this.MaxKey = key
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	this.Store[key]++
	this.balance(key)

}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	this.Store[key]--
	this.balance(key)
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if len(this.Store) != 0 {
		return this.MaxKey
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if len(this.Store) != 0 {
		return this.MinKey
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

