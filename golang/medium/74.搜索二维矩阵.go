/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.63%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    17.1K
 * Total Submissions: 47.6K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 * 示例 1:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * 输出: false
 *
 */

func bserach(nums []int, target int) bool {
	for l, r := 0, len(nums)-1; l <= r; {
		p := (l + r) / 2
		if nums[p] == target {
			return true
		} else {
			if nums[p] < target {
				l = p + 1
			} else {
				r = p - 1
			}
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	if n == 0 {
		return false
	}

	for l, r := 0, m-1; l <= r; {
		p := (l + r) / 2
		if matrix[p][0] <= target && matrix[p][n-1] >= target {
			return bserach(matrix[p], target)
		} else {
			if matrix[p][n-1] <= target {
				l = p + 1
			} else {
				r = p - 1
			}
		}
	}
	return false
}

