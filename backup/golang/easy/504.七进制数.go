package easy

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 *
 * https://leetcode-cn.com/problems/base-7/description/
 *
 * algorithms
 * Easy (45.59%)
 * Likes:    22
 * Dislikes: 0
 * Total Accepted:    5.5K
 * Total Submissions: 12K
 * Testcase Example:  '100'
 *
 * 给定一个整数，将其转化为7进制，并以字符串形式输出。
 *
 * 示例 1:
 *
 *
 * 输入: 100
 * 输出: "202"
 *
 *
 * 示例 2:
 *
 *
 * 输入: -7
 * 输出: "-10"
 *
 *
 * 注意: 输入范围是 [-1e7, 1e7] 。
 *
 */
import "strconv"

func convertToBase7(num int) string {
	// 不断除7 注意正负号
	if num == 0 {
		return strconv.Itoa(num)
	}
	needSign := false
	if num < 0 {
		num = num * -1
		needSign = true
	}

	res := ""
	for num > 0 {
		res = strconv.Itoa(num%7) + res
		num /= 7
	}
	if needSign {
		res = "-" + res
	}
	return res
}
