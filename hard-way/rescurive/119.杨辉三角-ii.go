/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (57.27%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    20.4K
 * Total Submissions: 35.6K
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

func getRow(rowIndex int) []int {
	// 递归
	if rowIndex == 0 {
		return []int{1}
	}
	lastRow := getRow(rowIndex - 1)
	row := make([]int, rowIndex+1)
	for i, _ := range row {
		if i == 0 || i == len(row)-1 {
			row[i] = 1
		} else {
			row[i] = lastRow[i-1] + lastRow[i]
		}
	}

	return row
}

func getRow(rowIndex int) []int {
	// 只保留上一行的数据

	last := []int{1}
	if rowIndex == 0 {
		return last
	}

	res := []int{}
	for i := 1; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			need := 1
			if 0 < j && j < i {
				need = last[j-1] + last[j]
			}
			res = append(res, need)
		}
		last = res
		res = []int{}
	}
	return last
}
