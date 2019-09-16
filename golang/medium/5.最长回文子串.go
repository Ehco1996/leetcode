/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (26.38%)
 * Likes:    1278
 * Dislikes: 0
 * Total Accepted:    108.3K
 * Total Submissions: 400K
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * 示例 1：
 *
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 *
 * 示例 2：
 *
 * 输入: "cbbd"
 * 输出: "bb"
 *
 *
 */
func longestPalindrome(s string) string {
	//https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zhong-xin-kuo-san-dong-tai-gui-hua-by-liweiwei1419/
	size := len(s)
	if size <= 1 {
		return s
	}

	//二维 dp 问题
	//状态：dp[l,r]: s[l:r] 包括 l，r ，表示的字符串是不是回文串
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size)
	}

	maxLens := 1
	res := string(s[0])

	//因为只有 1 个字符的情况在最开始做了判断
	//左边界一定要比右边界小，因此右边界从 1 开始
	for r := 1; r < size; r++ {
		for l := 0; l < r; l++ {
			//状态转移方程：如果头尾字符相等并且中间也是回文
			//在头尾字符相等的前提下，如果收缩以后不构成区间（最多只有 1 个元素），直接返回 True 即可
			//否则要继续看收缩以后的区间的回文性
			//重点理解 or 的短路性质在这里的作用
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if curLens := r - l + 1; curLens > maxLens {
					maxLens = curLens
					res = s[l : r+1]
				}
			}
		}
	}
	return res
}


