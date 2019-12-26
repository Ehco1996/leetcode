/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 *
 * https://leetcode-cn.com/problems/01-matrix/description/
 *
 * algorithms
 * Medium (36.73%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    8K
 * Total Submissions: 21.7K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
 *
 * 两个相邻元素间的距离为 1 。
 *
 * 示例 1:
 * 输入:
 *
 *
 * 0 0 0
 * 0 1 0
 * 0 0 0
 *
 *
 * 输出:
 *
 *
 * 0 0 0
 * 0 1 0
 * 0 0 0
 *
 *
 * 示例 2:
 * 输入:
 *
 *
 * 0 0 0
 * 0 1 0
 * 1 1 1
 *
 *
 * 输出:
 *
 *
 * 0 0 0
 * 0 1 0
 * 1 2 1
 *
 *
 * 注意:
 *
 *
 * 给定矩阵的元素个数不超过 10000。
 * 给定矩阵中至少有一个元素是 0。
 * 矩阵中的元素只在四个方向上相邻: 上、下、左、右。
 *
 *
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {
	// bfs
	xSize := len(matrix)
	if xSize == 0 {
		return matrix
	}
	ySize := len(matrix[0])

	res := make([][]int, xSize)
	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, ySize)
	}

	directions := [][]int{
		[]int{-1, 0},
		[]int{0, -1},
		[]int{1, 0},
		[]int{0, 1},
	}

	var find func(x, y int)
	find = func(x, y int) {
		queue := [][]int{
			[]int{x, y, 0},
		}
		for len(queue) > 0 {
			now := queue[0]
			nx, ny, cur := now[0], now[1], now[2]
			queue = queue[1:]
			if matrix[nx][ny] == 0 {
				res[x][y] = cur
				return
			}
			for _, dir := range directions {
				newX := nx + dir[0]
				newY := ny + dir[1]
				if newX >= 0 && newY >= 0 && newX < xSize && newY < ySize {
					queue = append(queue, []int{newX, newY, cur + 1})
				}
			}
		}
	}
	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			find(i, j)
		}
	}
	return res
}

// @lc code=end

