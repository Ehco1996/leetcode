/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 *
 * https://leetcode-cn.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (53.73%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    18.7K
 * Total Submissions: 34.6K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [1,1,1],
 * [1,0,1],
 * [1,1,1]
 * ]
 * 输出:
 * [
 * [1,0,1],
 * [0,0,0],
 * [1,0,1]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * [0,1,2,0],
 * [3,4,5,2],
 * [1,3,1,5]
 * ]
 * 输出:
 * [
 * [0,0,0,0],
 * [0,4,5,0],
 * [0,3,1,0]
 * ]
 *
 * 进阶:
 *
 *
 * 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个常数空间的解决方案吗？
 *
 *
 */

func setZero(matrix [][]int, m, n int) {
	for idx, _ := range matrix[m] {
		matrix[m][idx] = 0
	}

	for _, nums := range matrix {
		nums[n] = 0
	}
}

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := 0
	if m != 0 {
		n = len(matrix[0])
	}

	memory := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				memory = append(memory, []int{i, j})
			}
		}
	}

	for _, mem := range memory {
		setZero(matrix, mem[0], mem[1])
	}
}

