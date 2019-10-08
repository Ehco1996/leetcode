/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (42.78%)
 * Likes:    318
 * Dislikes: 0
 * Total Accepted:    30.8K
 * Total Submissions: 71.2K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *

 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */

// @lc code=start

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	// 使用数组 cell 保存每步子问题的最优解。
	// cell[i] 代表含第 i 个元素的最长上升子序列的长度。
	// 求解 cell[i] 时，向前遍历找出比 i 元素小的元素 j，令 cell[i] 为 max（cell[i],cell[j]+1)。
	if len(nums) == 0 {
		return 0
	}

	cell := []int{1}
	for i := 1; i < len(nums); i++ {
		cell = append(cell, 1)
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				cell[i] = max(cell[i], cell[j]+1)
			}
		}
	}
	sort.Ints(cell)
	return cell[len(cell)-1]
}

// @lc code=end

