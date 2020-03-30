/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 *
 * https://leetcode-cn.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (46.51%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    29.4K
 * Total Submissions: 62.9K
 * Testcase Example:  '1'
 *
 * 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: true
 * 解释: 2^0 = 1
 *
 * 示例 2:
 *
 * 输入: 16
 * 输出: true
 * 解释: 2^4 = 16
 *
 * 示例 3:
 *
 * 输入: 218
 * 输出: false
 *
 */
func isPowerOfTwo(n int) bool {
	// 不停的x2 看是否相等
	if n == 1 {
		return true
	}
	m := 1
	for m < n {
		m *= 2
	}
	return m == n

}
