/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 *
 * https://leetcode-cn.com/problems/longest-consecutive-sequence/description/
 *
 * algorithms
 * Hard (45.68%)
 * Likes:    170
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 34.6K
 * Testcase Example:  '[100,4,200,1,3,2]'
 *
 * 给定一个未排序的整数数组，找出最长连续序列的长度。
 *
 * 要求算法的时间复杂度为 O(n)。
 *
 * 示例:
 *
 * 输入: [100, 4, 200, 1, 3, 2]
 * 输出: 4
 * 解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
 *
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	// hash 存出现过的数字

	h := make(map[int]bool)
	ret := 0

	for _, nums := range nums {
		h[nums] = true
	}

	for num, _ := range h {
		// 如果num-1存在 那么一定被计算过了
		if !h[num-1] {
			cur := num
			lens := 1
			for h[cur+1] {
				cur = cur + 1
				lens++
			}
			if lens > ret {
				ret = lens
			}
		}
	}
	return ret
}

// @lc code=end

