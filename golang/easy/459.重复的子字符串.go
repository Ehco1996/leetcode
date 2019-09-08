// package golang

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (42.23%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 19.4K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 *
 * 示例 1:
 *
 *
 * 输入: "abab"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aba"
 *
 * 输出: False
 *
 *
 * 示例 3:
 *
 *
 * 输入: "abcabcabcabc"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 *
 *
 */
 func repeatedSubstringPattern(s string) bool {
	 // 滑动窗口
	if len(s) < 2 {
		return false
	}

	for r := 1; r <= len(s)/2; r++ {

		p := s[0:r]
		start := 0
		end := r

		// check
		for s[start:end] == p && (end+len(p)) <= len(s) {
			start += len(p)
			end += len(p)
		}

		if end == len(s)&& s[start:end] == p{
			return true
		}
	}
	return false
}
