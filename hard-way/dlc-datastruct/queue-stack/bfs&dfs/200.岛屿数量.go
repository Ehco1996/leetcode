/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 *
 * https://leetcode-cn.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (46.08%)
 * Likes:    286
 * Dislikes: 0
 * Total Accepted:    40.3K
 * Total Submissions: 87.5K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * 给定一个由 '1'（陆地）和
 * '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
 *
 * 示例 1:
 *
 * 输入:
 * 11110
 * 11010
 * 11000
 * 00000
 *
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入:
 * 11000
 * 11000
 * 00100
 * 00011
 *
 * 输出: 3
 *
 *
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	// bfs

	if len(grid) == 0 {
		return 0
	}
	X := len(grid)
	Y := len(grid[0])
	res := 0

	directions := [][]int{
		[]int{-1, 0},
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
	}

	for x := 0; x < X; x++ {
		for y := 0; y < Y; y++ {
			if grid[x][y] == '1' {
				// 写为0标记已经访问过了
				grid[x][y] = '0'
				res++
				queue := [][]int{
					[]int{x, y},
				}
				for len(queue) > 0 {
					cur_x, cur_y := queue[0][0], queue[0][1]
					queue = queue[1:]
					for _, dir := range directions {
						new_x := cur_x + dir[0]
						new_y := cur_y + dir[1]
						// 如果不越界，没访问过 并且还是陆地，就接着放入队列
						if 0 <= new_x && new_x < X &&
							0 <= new_y && new_y < Y &&
							grid[new_x][new_y] == '1' {
							queue = append(queue, []int{new_x, new_y})
							grid[new_x][new_y] = '0'
						}
					}
				}
			}
		}
	}
	return res
}

// @lc code=end

