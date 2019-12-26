package medium

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (71.28%)
 * Likes:    532
 * Dislikes: 0
 * Total Accepted:    42.1K
 * Total Submissions: 58.6K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 *
 * 例如，给出 n = 3，生成结果为：
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 *
 */
func generateParenthesis(n int) []string {
	// dfs
	h := make(map[string]int)
	var dfs func(path string)
	dfs = func(path string) {
		if len(path) == 2*n {
			h[path]++
			return
		}
		for idx, _ := range path {
			dfs(path[:idx] + "()" + path[idx:])
		}
	}

	res := []string{}
	dfs("()")
	for p, _ := range h {
		res = append(res, p)
	}
	return res

}
