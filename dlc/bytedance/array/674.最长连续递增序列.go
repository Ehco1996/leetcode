/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 *
 * https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (43.14%)
 * Likes:    59
 * Dislikes: 0
 * Total Accepted:    16.1K
 * Total Submissions: 36.9K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * 给定一个未经排序的整数数组，找到最长且连续的的递增序列。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,5,4,7]
 * 输出: 3
 * 解释: 最长连续递增序列是 [1,3,5], 长度为3。
 * 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,2,2,2,2]
 * 输出: 1
 * 解释: 最长连续递增序列是 [2], 长度为1。
 *
 *
 * 注意：数组长度不会超过10000。
 *
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	res := 0
	cur := 0

	for idx, num := range nums {
		if idx > 0 {
			if nums[idx-1] < num {
				cur++
			} else {
				cur = 1
			}
		} else {
			cur++
		}
		if cur > res {
			res = cur
		}
	}
	return res
}

// @lc code=end

