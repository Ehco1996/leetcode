/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (63.19%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    13.8K
 * Total Submissions: 21.6K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 *
 * 返回 s 所有可能的分割方案。
 *
 * 示例:
 *
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 */

// @lc code=start
func partition(s string) [][]string {
	// dp + dfs dp缓存结果 dp[i][j]表示s[i:j]为回文串

	size := len(s)
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
	}

	for r := 0; r < size; r++ {
		for l := 0; l < r+1; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			}
		}
	}

	res := [][]string{}
	var dfs func(start int, temp []string)
	dfs = func(start int, temp []string) {
		if start == size {
			//如果分割完了
			res = append(res, temp)
		}
		for j := start; j < size; j++ {
			if dp[start][j] {
				temp = append(temp, s[start:j+1])
				dfs(j+1, temp)
				temp = temp[:len(temp)-1]
			}
		}
	}
	t := []string{}
	dfs(0, t)
	return res
}

// @lc code=end