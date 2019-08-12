package main

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (63.35%)
 * Likes:    174
 * Dislikes: 0
 * Total Accepted:    32.1K
 * Total Submissions: 50.6K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
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

// func main() {
// 	fmt.Println(generate(4))

// }
