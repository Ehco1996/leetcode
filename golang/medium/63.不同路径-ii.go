/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (31.69%)
 * Likes:    165
 * Dislikes: 0
 * Total Accepted:    23K
 * Total Submissions: 72.3K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// d(m,n) = d(m,n-1) if(d(m,n-1)!=0) + d(m-1,n) if(d(m-1,n)=0)

	if len(obstacleGrid) == 0 {
		return 1
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 检查第一列
	ybug := false
	for idx, nums := range dp {
		if obstacleGrid[idx][0] != 1 && !ybug {
			nums[0] = 1
		} else {
			ybug = true
			nums[0] = -1
		}
	}

	// 检查第一行
	xbug := false
	for idx, _ := range dp[0] {
		if obstacleGrid[0][idx] != 1 && !xbug {
			dp[0][idx] = 1
		} else {
			xbug = true
			dp[0][idx] = -1
		}
	}

	var f func(x, y int) int
	f = func(x, y int) int {
		if dp[x][y] == 0 {
			tmp := 0
			if obstacleGrid[x-1][y] != 1 {
				tmp += f(x-1, y)
			}
			if obstacleGrid[x][y-1] != 1 {
				tmp += f(x, y-1)
			}
			dp[x][y] = tmp
			return tmp
		} else if dp[x][y] == -1 {
			return 0
		}
		return dp[x][y]
	}
	return f(m-1, n-1)
}

