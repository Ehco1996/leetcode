/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (58.63%)
 * Likes:    91
 * Dislikes: 0
 * Total Accepted:    28.8K
 * Total Submissions: 49K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */

// @lc code=start
func getRow(rowIndex int) []int {
	//只保留上一行

	lastRow := []int{1}
	if rowIndex == 0 {
		return lastRow
	}

	for i := 2; i <= rowIndex+1; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j < i-1; j++ {
			row[j] = lastRow[j-1] + lastRow[j]
		}
		lastRow = row
	}
	return lastRow
}

// @lc code=end

