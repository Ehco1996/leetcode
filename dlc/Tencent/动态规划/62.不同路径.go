/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (55.20%)
 * Likes:    326
 * Dislikes: 0
 * Total Accepted:    40.9K
 * Total Submissions: 72.6K
 * Testcase Example:  '3\n2'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 *
 *
 * 示例 2:
 *
 * 输入: m = 7, n = 3
 * 输出: 28
 *
 */
func uniquePaths(m int, n int) int {
	// 二维dp
	// d(m,n) = d(m,n-1) + d(m-1,n)
	if m <= 1 || n <= 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for _, nums := range dp {
		nums[0] = 1
	}
	for idx, _ := range dp[0] {
		dp[0][idx] = 1

	}

	var f func(x, y int) int
	f = func(x, y int) int {
		if dp[x][y] == 0 {
			// 这个格子还没来来过
			dp[x][y] = f(x-1, y) + f(x, y-1)
		}
		return dp[x][y]
	}
	return f(m-1, n-1)
}
