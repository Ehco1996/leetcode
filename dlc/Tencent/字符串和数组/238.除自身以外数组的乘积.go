/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 *
 * https://leetcode-cn.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (62.21%)
 * Likes:    218
 * Dislikes: 0
 * Total Accepted:    18.2K
 * Total Submissions: 28.7K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i]
 * 之外其余各元素的乘积。
 *
 * 示例:
 *
 * 输入: [1,2,3,4]
 * 输出: [24,12,8,6]
 *
 * 说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
 *
 * 进阶：
 * 你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
 *
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	// 双指针 先从左面累乘 在从右面累乘
	res := make([]int, len(nums))
	res[0] = 1
	for idx := 1; idx < len(nums); idx++ {
		res[idx] = res[idx-1] * nums[idx-1]
	}

	right := 1
	for idx := len(nums) - 1; idx > -1; idx-- {
		res[idx] *= right
		right *= nums[idx]
	}
	return res
}

// @lc code=end

