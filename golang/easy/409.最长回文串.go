package easy

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 *
 * https://leetcode-cn.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (49.39%)
 * Likes:    65
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 18.2K
 * Testcase Example:  '"abccccdd"'
 *
 * 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
 *
 * 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
 *
 * 注意:
 * 假设字符串的长度不会超过 1010。
 *
 * 示例 1:
 *
 *
 * 输入:
 * "abccccdd"
 *
 * 输出:
 * 7
 *
 * 解释:
 * 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 *
 *
 */
func longestPalindrome(s string) int {
	// hash 是串不是子串
	// 勾陈回文串的条件  双数 / 单数中取最大双数 / 最大的单数对可以放串中间
	h := make(map[rune]int)
	for _, c := range s {
		h[c] += 1
	}
	c := 0
	fmax := 0
	for _, v := range h {
		if v%2 == 0 {
			c += v
		} else if v > fmax {
			c += fmax / 2 * 2
			fmax = v
		} else {
			c += v / 2 * 2
		}
	}

	return c + fmax

}
