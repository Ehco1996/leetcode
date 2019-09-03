// package main

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 *
 * https://leetcode-cn.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (39.21%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    17.4K
 * Total Submissions: 44.2K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，返回 n! 结果尾数中零的数量。
 *
 * 示例 1:
 *
 * 输入: 3
 * 输出: 0
 * 解释: 3! = 6, 尾数中没有零。
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: 1
 * 解释: 5! = 120, 尾数中有 1 个零.
 *
 * 说明: 你算法的时间复杂度应为 O(log n) 。
 *
 */
func trailingZeroes(n int) int {
	//因为是阶乘后的0
	//0是由于2*5得出一个0
	//因此便是求2*5有多少个
	//因为2<5
	//有5一定有2
	//因此便是求5的个数
	c := 0
	for n >= 5 {
		c += n / 5
		n /= 5
	}
	return c
}
