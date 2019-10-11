/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 *
 * https://leetcode-cn.com/problems/shuffle-an-array/description/
 *
 * algorithms
 * Medium (46.66%)
 * Likes:    35
 * Dislikes: 0
 * Total Accepted:    10.7K
 * Total Submissions: 22.4K
 * Testcase Example:  '["Solution","shuffle","reset","shuffle"]\n[[[1,2,3]],[],[],[]]'
 *
 * 打乱一个没有重复元素的数组。
 *
 * 示例:
 *
 *
 * // 以数字集合 1, 2 和 3 初始化数组。
 * int[] nums = {1,2,3};
 * Solution solution = new Solution(nums);
 *
 * // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
 * solution.shuffle();
 *
 * // 重设数组到它的初始状态[1,2,3]。
 * solution.reset();
 *
 * // 随机返回数组[1,2,3]打乱后的结果。
 * solution.shuffle();
 *
 *
 */

// @lc code=start
import "math/rand"

type Solution struct {
	a []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.a
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	// 随机交换元素的位置
	tmp := make([]int, len(this.a))
	copy(tmp, this.a)
	for i := 0; i < len(tmp); i++ {
		r := rand.Intn(len(tmp))
		tmp[i], tmp[r] = tmp[r], tmp[i]
	}
	return tmp
}

// @lc code=end

