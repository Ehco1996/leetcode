/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 *
 * https://leetcode-cn.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (48.13%)
 * Likes:    1623
 * Dislikes: 0
 * Total Accepted:    160.6K
 * Total Submissions: 326.4K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 *
 * 示例:
 *
 * 输入: [-2,1,-3,4,-1,2,1,-5,4],
 * 输出: 6
 * 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
 *
 *
 * 进阶:
 *
 * 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 *
 */

// @lc code=start
func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSubArray(nums []int) int {
	// dp(i) = max(dp[i-1]+now,now)
	if len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	cur := nums[0]
	for _, num := range nums[1:] {
		cur = max(cur+num, num)
		sum = max(sum, cur)
	}
	return sum
}

// @lc code=end

