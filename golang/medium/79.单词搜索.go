/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (38.50%)
 * Likes:    213
 * Dislikes: 0
 * Total Accepted:    20.4K
 * Total Submissions: 52.2K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 * 示例:
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true.
 * 给定 word = "SEE", 返回 true.
 * 给定 word = "ABCB", 返回 false.
 *
 */
func exist(board [][]byte, word string) bool {

	m := len(board)
	if m == 0 {
		return false
	}
	n := 0
	if m > 0 {
		n = len(board[0])
	}

	// 使用标记位置
	mark := make([][]int, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]int, n)
	}

	var helper func(x, y int, target string) bool
	helper = func(x, y int, target string) bool {
		if len(target) == 0 {
			return true
		}
		if x > m-1 || y > n-1 || x < 0 || y < 0 || mark[x][y] == 1 {
			return false
		}
		if board[x][y] == byte(target[0]) {
			target = target[1:]
			mark[x][y] = 1
			res := helper(x-1, y, target) || helper(x+1, y, target) || helper(x, y+1, target) || helper(x, y-1, target)
			if !res {
				mark[x][y] = 0
			}
			return res
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helper(i, j, word) {
				return true
			}
		}
	}
	return false
}

