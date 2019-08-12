package main

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

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	res := [][]int{{}}
	for i := 0; i < numRows; i++ {
		if i > 0 {
			res = append(res, []int{})
		}
		for j := 0; j <= i; j++ {
			need := 1
			if 0 < j && j < i {
				need = res[i-1][j-1] + res[i-1][j]
			}
			res[i] = append(res[i], need)
		}
	}
	return res
}
func getRow(rowIndex int) []int {

	triangle := generate(rowIndex + 1)
	return triangle[rowIndex]
}

// func main() {
// 	fmt.Println(getRow(0))
// }
