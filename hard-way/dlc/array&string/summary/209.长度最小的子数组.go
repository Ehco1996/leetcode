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
func minSubArrayLen(s int, nums []int) int {
	// 双指针+滑动窗口
	if len(nums) == 0 {
		return 0
	}
	res := len(nums) + 1
	cur := 0
	for i, j := 0, 0; j < len(nums); j++ {
		cur += nums[j]
		if cur >= s {
			//开始移动左指针
			for cur-nums[i] >= s {
				cur -= nums[i]
				i += 1
			}
			if j-i+1 <= res {
				res = j - i + 1
			}
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}

// @lc code=end

