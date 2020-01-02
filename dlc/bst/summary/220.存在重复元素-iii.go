/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode-cn.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (25.45%)
 * Likes:    118
 * Dislikes: 0
 * Total Accepted:    10.9K
 * Total Submissions: 42.8K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j
 * 之间的差的绝对值最大为 ķ。
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3, t = 0
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1, t = 2
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,5,9,1,5,9], k = 2, t = 3
 * 输出: false
 *
 */

// @lc code=start

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i, a := range nums {
		for j, b := range nums {
			if i != j && abs(i,j) <= k && abs(a,b) <= t {
				return true
			}
		}
	}
	return false
}

// @lc code=end

