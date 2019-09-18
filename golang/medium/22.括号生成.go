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
	// n = 1  ()
	// n = 2  ()() / (())
	// n = 3  ()()() / (())() / ()(()) / ((())) / (()())/

	last := []string{"()"}
	for i := 2; i <= n; i++ {
		temp := []string{}
		for _, now := range last {
			for idx := 0; idx <= len(now); idx++ {
				temp = append(temp, now[:idx]+"()"+now[idx:])
			}
		}
		last = temp
	}
	res := []string{}
	h := make(map[string]int)
	for _, pad := range last {
		if _, ok := h[pad]; !ok {
			res = append(res, pad)
			h[pad]++
		}
	}
	return res

}

