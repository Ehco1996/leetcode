/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (36.87%)
 * Likes:    210
 * Dislikes: 0
 * Total Accepted:    24.4K
 * Total Submissions: 65.9K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */
func reverse(nums []int) {
	// 翻转
	for l, r := 0, len(nums)-1; l < r; {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func transpose(matrix [][]int) [][]int {
	// 转置 A[i,j] -> A[j,i]
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}

	if m*n <= 1 {
		return matrix
	}

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}

func lrotate(matrix [][]int) [][]int {
	// 顺时针旋转矩阵
	// 先上下翻转，再转置。
	for i := 0; i < len(matrix); i++ {
		reverse(matrix[i])
	}
	return transpose(matrix)
}

func spiralOrder(matrix [][]int) []int {
	// 拿矩阵的第一行，然后逆时针旋转剩下的矩阵
	res := []int{}
	for len(matrix) > 0 {
		res = append(res, matrix[0]...)
		matrix = lrotate(matrix[1:])
	}
	return res
}

