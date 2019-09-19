package easy

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Easy (38.20%)
 * Likes:    151
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 24.3K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 *
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 *
 * 说明：
 *
 *
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 *
 *
 * 示例 1:
 *
 *
 * 输入:
 * s: "cbaebabacd" p: "abc"
 *
 * 输出:
 * [0, 6]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s: "abab" p: "ab"
 *
 * 输出:
 * [0, 1, 2]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 *
 *
 */

func CopyMap(m map[rune]int) map[rune]int {
	cp := make(map[rune]int)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func helper(s string, h map[rune]int) bool {

	for _, r := range s {
		v, ok := h[r]
		if !ok || v < 1 {
			return false
		} else {
			h[r]--
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	// 滑动窗口 + hash判段
	h := make(map[rune]int)
	for _, r := range p {
		h[r]++
	}

	res := []int{}
	lens := len(p)
	for i := 0; i <= len(s)-lens; i++ {
		if helper(s[i:i+lens], CopyMap(h)) {
			res = append(res, i)
		}
	}
	return res
}
