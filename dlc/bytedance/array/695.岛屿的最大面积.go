/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 *
 * https://leetcode-cn.com/problems/max-area-of-island/description/
 *
 * algorithms
 * Medium (55.61%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    16.1K
 * Total Submissions: 27.3K
 * Testcase Example:  '[[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]'
 *
 * 给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地)
 * 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
 *
 * 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
 *
 * 示例 1:
 *
 *
 * [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,1,1,0,1,0,0,0,0,0,0,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,0,1,0,0],
 * ⁠[0,1,0,0,1,1,0,0,1,1,1,0,0],
 * ⁠[0,0,0,0,0,0,0,0,0,0,1,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,1,0,0,0],
 * ⁠[0,0,0,0,0,0,0,1,1,0,0,0,0]]
 *
 *
 * 对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
 *
 * 示例 2:
 *
 *
 * [[0,0,0,0,0,0,0,0]]
 *
 * 对于上面这个给定的矩阵, 返回 0。
 *
 * 注意: 给定的矩阵grid 的长度和宽度都不超过 50。
 *
 */

// @lc code=start

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func find(grid [][]int, x, y int) int {
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[0])-1 || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	res := 1
	res += find(grid, x+1, y)
	res += find(grid, x-1, y)
	res += find(grid, x, y+1)
	res += find(grid, x, y-1)
	return res
}

func maxAreaOfIsland(grid [][]int) int {
	// dfs 上下左右找，要把找到的地方标记为0
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res = max(res, find(grid, i, j))
		}
	}
	return res
}

// @lc code=end

