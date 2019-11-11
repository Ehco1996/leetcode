/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (40.05%)
 * Likes:    145
 * Dislikes: 0
 * Total Accepted:    19.3K
 * Total Submissions: 48K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
 *
 * 示例:
 *
 * 输入: s = 7, nums = [2,3,1,2,4,3]
 * 输出: 2
 * 解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
 *
 *
 * 进阶:
 *
 * 如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 */

// @lc code=start
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	// 双指针+滑动窗口
	ret := len(nums) + 1
	sum := 0

	for l, r := 0, 0; r < len(nums); {
		for sum < s && r < len(nums) {
			sum += nums[r]
			r++
		}
		for sum >= s&&l >= 0 {
			ret = min(ret, r-l)
			sum -= nums[l]
			l++
		}
	}
	if ret == len(nums)+1 {
		return 0
	}
	return ret
}

// @lc code=end

