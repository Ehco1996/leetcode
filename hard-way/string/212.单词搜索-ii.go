/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (35.29%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    4.9K
 * Total Submissions: 13.5K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 * 示例:
 *
 * 输入:
 * words = ["oath","pea","eat","rain"] and board =
 * [
 * ⁠ ['o','a','a','n'],
 * ⁠ ['e','t','a','e'],
 * ⁠ ['i','h','k','r'],
 * ⁠ ['i','f','l','v']
 * ]
 *
 * 输出: ["eat","oath"]
 *
 * 说明:
 * 你可以假设所有输入都由小写字母 a-z 组成。
 *
 * 提示:
 *
 *
 * 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
 * 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
 * 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
 *
 *
*/

// @lc code=start
func dfs(x, y int, target string, boardP *[][]byte, markP *[][]int) bool {
	if len(target) == 0 {
		return true
	}

	board := *boardP
	mark := *markP

	m := len(board)
	n := 0
	if m > 0 {
		n = len(board[0])
	}
	if x > m-1 || y > n-1 || x < 0 || y < 0 || mark[x][y] == 1 {
		return false
	}
	if board[x][y] == byte(target[0]) {
		target = target[1:]
		mark[x][y] = 1
		res := dfs(x-1, y, target, boardP, markP) ||
			dfs(x+1, y, target, boardP, markP) ||
			dfs(x, y+1, target, boardP, markP) ||
			dfs(x, y-1, target, boardP, markP)
		if !res {
			mark[x][y] = 0
		}
		return res
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	// 深度优先搜索

	m := len(board)
	n := 0
	if m > 0 {
		n = len(board[0])
	}

	res := []string{}

	for _, word := range words {
		// 使用标记位置
		mark := make([][]int, m)
		for i := 0; i < m; i++ {
			mark[i] = make([]int, n)
		}
		found := false
		for i := 0; i < m && !found; i++ {
			for j := 0; j < n && !found; j++ {
				if dfs(i, j, word, &board, &mark) {
					res = append(res, word)
					found = true
				}
			}
		}
	}
	return res
}

// @lc code=end

