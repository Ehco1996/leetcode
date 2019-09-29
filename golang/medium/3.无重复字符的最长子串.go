// package medium

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (30.62%)
 * Likes:    2391
 * Dislikes: 0
 * Total Accepted:    212.9K
 * Total Submissions: 679.5K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口 + hash 来确定改字母是否访问过
	if len(s) <= 1 {
		return len(s)
	}

	res := 0
	h := make(map[byte]int)

	for left, right := 0, 0; left <= right && right < len(s); {
		v, ok := h[s[right]]
		if v == 0 || !ok {
			h[s[right]]++
			right++
			res = max(res, len(s[left:right]))
		} else {
			// 左边窗口右移 直到满足要求
			h[s[left]]--
			left++
		}
	}
	return res
}
