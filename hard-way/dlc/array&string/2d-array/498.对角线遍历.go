/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 *
 * https://leetcode-cn.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (38.86%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    7.8K
 * Total Submissions: 20.1K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
 *
 *
 *
 * 示例:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 *
 * 输出:  [1,2,4,7,5,3,6,8,9]
 *
 * 解释:
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 * 给定矩阵中的元素总数不会超过 100000 。
 *
 *
 */

// @lc code=start

// class Solution:
//     def findDiagonalOrder(self, matrix: List[List[int]]) -> List[int]:
//         if len(matrix) == 0:
//             return []
//         M, N, result = len(matrix), len(matrix[0]), []
//         for curve_line in range(M + N - 1):
//             row_begin = 0 if curve_line + 1 <= N else curve_line + 1 - N
//             row_end = M - 1 if curve_line + 1 >= M else curve_line
//             if curve_line % 2 == 1:
//                 for i in range(row_begin,row_end + 1):
//                     result.append(matrix[i][curve_line - i])
//             else:
//                 for i in range(row_end,row_begin - 1,-1):
//                     result.append(matrix[i][curve_line - i])
//         return result

func findDiagonalOrder(matrix [][]int) []int {
	// 找规律
	res := []int{}
	if len(matrix) == 0 {
		return res
	}

	M, N := len(matrix), len(matrix[0])
	for curve_line := 0; curve_line < M+N-1; curve_line++ {
		begin := 0
		if curve_line+1 > N {
			begin = curve_line + 1 - N
		}
		end := curve_line
		if curve_line+1 > M {
			end = M - 1
		}

		if curve_line%2 == 1 {
			for i := begin; i < end+1; i++ {
				res = append(res, matrix[i][curve_line-i])
			}
		} else {
			for i := end; i > begin-1; i-- {
				res = append(res, matrix[i][curve_line-i])
			}
		}
	}
	return res
}

// @lc code=end

