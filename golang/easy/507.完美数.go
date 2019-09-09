/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 *
 * https://leetcode-cn.com/problems/perfect-number/description/
 *
 * algorithms
 * Easy (35.30%)
 * Likes:    35
 * Dislikes: 0
 * Total Accepted:    7K
 * Total Submissions: 19.7K
 * Testcase Example:  '28'
 *
 * 对于一个 正整数，如果它和除了它自身以外的所有正因子之和相等，我们称它为“完美数”。
 *
 * 给定一个 整数 n， 如果他是完美数，返回 True，否则返回 False
 *
 *
 *
 * 示例：
 *
 * 输入: 28
 * 输出: True
 * 解释: 28 = 1 + 2 + 4 + 7 + 14
 *
 *
 *
 *
 * 提示：
 *
 * 输入的数字 n 不会超过 100,000,000. (1e8)
 *
 */
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	l := 1
	r := num

	for l <= r && l*r == num {
		last := l
		for add := 1; l+add < r; add++ {
			if num%(l+add) == 0 {
				l += add
				r = num / l
				break
			}
		}
		if l == last || sum > num {
			break
		} else {
			sum = sum + l + r
		}
	}
	return sum == num
}

