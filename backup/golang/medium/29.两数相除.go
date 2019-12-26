package medium

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (18.52%)
 * Likes:    182
 * Dislikes: 0
 * Total Accepted:    21K
 * Total Submissions: 112.6K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 *
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 *
 * 示例 1:
 *
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 *
 * 说明:
 *
 *
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 *
 *
 */
import "math"

func toPositive(num int) int {
	if num >= 0 {
		return num
	}
	t := num
	num -= t
	num -= t
	return num
}

func toNegative(num int) int {
	// 除法变成减法
	if num < 0 {
		return num
	}
	t := num
	num -= t
	num -= t
	return num
}
func divide(dividend int, divisor int) int {
	res := 0
	flag := 1

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		flag = -1
	}

	dividend = toPositive(dividend)
	divisor = toPositive(divisor)

	for dividend >= divisor {
		res += 1
		dividend -= divisor
	}

	if flag == -1 {
		res = toNegative(res)
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
