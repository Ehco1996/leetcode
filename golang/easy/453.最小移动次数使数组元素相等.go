package easy

import "sort"

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Easy (51.85%)
 * Likes:    73
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 10K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。
 *
 * 示例:
 *
 *
 * 输入:
 * [1,2,3]
 *
 * 输出:
 * 3
 *
 * 解释:
 * 只需要3次移动（注意每次移动会增加两个元素的值）：
 *
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 *
 *
 */

func minMoves(nums []int) int {
	// 每个元素需要增加diff：now-min 次才能使得数组平衡
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	c := 0
	sort.Ints(nums)
	min := nums[0]
	for i := len(nums) - 1; i > 0; i-- {
		c += nums[i] - min
	}
	return c
}
