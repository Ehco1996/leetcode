package easy

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 *
 * https://leetcode-cn.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (46.15%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    12K
 * Total Submissions: 25.7K
 * Testcase Example:  '16'
 *
 * 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 16
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */
func isPowerOfFour(num int) bool {
	// 32 位以内最大的4次幂是 17179869184
	if num == 1 {
		return true
	}
	r := 1
	for r < num {
		r *= 4
		if r == num {
			return true
		}
	}
	return false
}
