/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (33.22%)
 * Likes:    1722
 * Dislikes: 0
 * Total Accepted:    288K
 * Total Submissions: 853.6K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */

// @lc code=start
import "math"
import "strconv"

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	flag := false
	if x < 0 {
		flag = true
		x *= -1
	}
	s := []byte(strconv.Itoa(x))
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	idx := 0
	for ; idx == '0'; idx++ {
	}
	n, _ := strconv.Atoi(string(s[idx:]))
	if flag {
		n *= -1
	}
	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}
	return n
}

// @lc code=end

