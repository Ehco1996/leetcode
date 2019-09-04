package main

import "fmt"

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 *
 * https://leetcode-cn.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (44.91%)
 * Likes:    65
 * Dislikes: 0
 * Total Accepted:    24.4K
 * Total Submissions: 54.1K
 * Testcase Example:  '27'
 *
 * 给定一个整数，写一个函数来判断它是否是 3 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 27
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 0
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: 9
 * 输出: true
 *
 * 示例 4:
 *
 * 输入: 45
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */
func isPowerOfThree(n int) bool {
	if n ==1{
		return true
	}
	r := 1
	for r < n {
		r *= 3
		if r == n {
			return true
		}
	}
	return false

}

// func main() {
// 	fmt.Println(isPowerOfThree(19684))
// }
