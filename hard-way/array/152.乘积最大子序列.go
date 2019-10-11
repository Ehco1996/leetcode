/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (35.25%)
 * Likes:    272
 * Dislikes: 0
 * Total Accepted:    21.2K
 * Total Submissions: 59.7K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */

// @lc code=start

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func maxProduct(nums []int) int {
	// dp 当num是负数的时候，最大值和最小值会发生转换
	// dp(max) = max(now,now*max)
	// dp(min) = min(now,now*max)

	if len(nums) == 0 {
		return 0
	}

	curMax := nums[0]
	curMin := nums[0]
	res := nums[0]

	for _, num := range nums[1:] {
		if num < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = max(curMax*num, num)
		curMin = min(curMin*num, num)
		res = max(res, curMax)
	}
	return res
}

// @lc code=end

