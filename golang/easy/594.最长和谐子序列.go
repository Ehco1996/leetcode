/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 *
 * https://leetcode-cn.com/problems/longest-harmonious-subsequence/description/
 *
 * algorithms
 * Easy (41.88%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    4.7K
 * Total Submissions: 11K
 * Testcase Example:  '[1,3,2,2,5,2,3,7]'
 *
 * 和谐数组是指一个数组里元素的最大值和最小值之间的差别正好是1。
 *
 * 现在，给定一个整数数组，你需要在所有可能的子序列中找到最长的和谐子序列的长度。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,2,2,5,2,3,7]
 * 输出: 5
 * 原因: 最长的和谐数组是：[3,2,2,2,3].
 *
 *
 * 说明: 输入的数组长度最大不超过20,000.
 *
 */
func findLHS(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	h := make(map[int]int)
	for _, c := range nums {
		h[c]++
	}

	cnt := 0
	for k, _ := range h {
		if _, ok := h[k+1]; ok && h[k+1]+h[k] > cnt {
			cnt = h[k+1] + h[k]
		}
	}
	return cnt
}

