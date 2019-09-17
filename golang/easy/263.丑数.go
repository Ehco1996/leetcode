package easy

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 *
 * https://leetcode-cn.com/problems/ugly-number/description/
 *
 * algorithms
 * Easy (47.12%)
 * Likes:    69
 * Dislikes: 0
 * Total Accepted:    14.7K
 * Total Submissions: 30.8K
 * Testcase Example:  '6'
 *
 * 编写一个程序判断给定的数是否为丑数。
 *
 * 丑数就是只包含质因数 2, 3, 5 的正整数。
 *
 * 示例 1:
 *
 * 输入: 6
 * 输出: true
 * 解释: 6 = 2 × 3
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: true
 * 解释: 8 = 2 × 2 × 2
 *
 *
 * 示例 3:
 *
 * 输入: 14
 * 输出: false
 * 解释: 14 不是丑数，因为它包含了另外一个质因数 7。
 *
 * 说明：
 *
 *
 * 1 是丑数。
 * 输入不会超过 32 位有符号整数的范围: [−2^31,  2^31 − 1]。
 *
 *
 */
func isUgly(num int) bool {
	// 暴力 不停的除这这三个因数
	unums := []int{5, 3, 2}
	for num != 1 {
		last := num
		for _, u := range unums {
			if num%u == 0 {
				num = num / u
				break
			}
		}
		if num == last {
			break
		}
	}
	if num == 1 {
		return true
	}
	if num != 5 && num != 3 && num != 2 {
		return false
	}
	return true
}
