/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (36.79%)
 * Likes:    126
 * Dislikes: 0
 * Total Accepted:    21K
 * Total Submissions: 56.2K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
 *
 *
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 *
 *
 * 示例:
 *
 * 现有矩阵 matrix 如下：
 *
 * [
 * ⁠ [1,   4,  7, 11, 15],
 * ⁠ [2,   5,  8, 12, 19],
 * ⁠ [3,   6,  9, 16, 22],
 * ⁠ [10, 13, 14, 17, 24],
 * ⁠ [18, 21, 23, 26, 30]
 * ]
 *
 *
 * 给定 target = 5，返回 true。
 *
 * 给定 target = 20，返回 false。
 *
*/

// @lc code=start

func bs(nums []int, target int) bool {
	for l, r := 0, len(nums)-1; l <= r; {
		p := (l + r) / 2
		if nums[p] == target {
			return true
		} else if nums[p] < target {
			l = p + 1
		} else {
			r = p - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	// 二分查找O(logN)
	res := false

	for _, nums := range matrix {
		if len(nums) == 0 {
			return false
		}
		if nums[0] <= target && nums[len(nums)-1] >= target {
			res = bs(nums, target)
		}
		if res {
			return res
		}
	}
	return res
}

// @lc code=end

