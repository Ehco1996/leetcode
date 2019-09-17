package easy

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (35.48%)
 * Likes:    104
 * Dislikes: 0
 * Total Accepted:    20.3K
 * Total Submissions: 56.8K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j
 * 的差的绝对值最大为 k。
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,2,3,1,2,3], k = 2
 * 输出: false
 *
 */

func containsNearbyDuplicate(nums []int, k int) bool {
	h := make(map[int][]int)
	for idx, num := range nums {
		h[num] = append(h[num], idx)

		if len(h[num]) > 1 {
			for _, lastIdx := range h[num] {
				if idx-lastIdx <= k && idx != lastIdx {
					return true
				}
			}
		}
	}
	return false
}
