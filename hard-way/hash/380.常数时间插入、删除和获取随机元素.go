/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 *
 * https://leetcode-cn.com/problems/insert-delete-getrandom-o1/description/
 *
 * algorithms
 * Medium (45.83%)
 * Likes:    51
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 11.8K
 * Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n' +
  '[[],[1],[2],[2],[],[1],[2],[]]'
 *
 * 设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
 *
 *
 * insert(val)：当元素 val 不存在时，向集合中插入该项。
 * remove(val)：元素 val 存在时，从集合中移除该项。
 * getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
 *
 *
 * 示例 :
 *
 *
 * // 初始化一个空的集合。
 * RandomizedSet randomSet = new RandomizedSet();
 *
 * // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
 * randomSet.insert(1);
 *
 * // 返回 false ，表示集合中不存在 2 。
 * randomSet.remove(2);
 *
 * // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
 * randomSet.insert(2);
 *
 * // getRandom 应随机返回 1 或 2 。
 * randomSet.getRandom();
 *
 * // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
 * randomSet.remove(1);
 *
 * // 2 已在集合中，所以返回 false 。
 * randomSet.insert(2);
 *
 * // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
 * randomSet.getRandom();
 *
 *
*/

// @lc code=start

type RandomizedSet struct {
	numsMap   map[int]int
	numsArray []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{numsArray: []int{}, numsMap: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numsMap[val]; ok {
		return false
	}

	this.numsMap[val] = len(this.numsArray)
	this.numsArray = append(this.numsArray, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	// 找到这个数的iindex并和最后一个元素互换
	index, ok := this.numsMap[val]
	if !ok {
		return false
	}

	last := this.numsArray[len(this.numsArray)-1]
	if val == last { // if we are removing the last element in the array, we don't need to swap and update index in map
		this.numsArray = this.numsArray[:len(this.numsArray)-1]
	} else {
		// swap and remove the last element makes the delete of the element O(1)
		this.numsArray[index] = last
		this.numsArray = this.numsArray[:len(this.numsArray)-1]
		// note: we also need to update the index of the last element in the map
		this.numsMap[this.numsArray[index]] = index
	}
	delete(this.numsMap, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.numsArray))

	return this.numsArray[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

