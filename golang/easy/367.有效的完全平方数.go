package main
import "fmt"
/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 *
 * https://leetcode-cn.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (41.67%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    14K
 * Total Submissions: 33.3K
 * Testcase Example:  '16'
 *
 * 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
 *
 * 说明：不要使用任何内置的库函数，如  sqrt。
 *
 * 示例 1：
 *
 * 输入：16
 * 输出：True
 *
 * 示例 2：
 *
 * 输入：14
 * 输出：False
 *
 *
 */
func isPerfectSquare(num int) bool {
	if num==1{
		return true
	}
	for n := num / 2; n*n >= num;n-=1 {
		if n*n == num{
			return true
		}
	}
	return false
}

// func main(){
// 	fmt.Println(isPerfectSquare(32))
// }