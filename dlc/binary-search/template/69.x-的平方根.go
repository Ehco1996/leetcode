/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (37.22%)
 * Likes:    267
 * Dislikes: 0
 * Total Accepted:    77.5K
 * Total Submissions: 208.2K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l, r := 0, x
	var mid int
	var sqr int

	for {
		mid = (l + r) / 2
		if mid == l {
			return mid
		}

		sqr = mid * mid
		if sqr == x {
			return mid
		}
		if sqr > x {
			r = mid
		}
		if sqr < x {
			l = mid
		}
	}
}

// @lc code=end

