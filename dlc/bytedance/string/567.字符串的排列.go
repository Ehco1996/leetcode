/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (33.04%)
 * Likes:    82
 * Dislikes: 0
 * Total Accepted:    13.7K
 * Total Submissions: 41.1K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 *
 * 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 *
 * 示例1:
 *
 *
 * 输入: s1 = "ab" s2 = "eidbaooo"
 * 输出: True
 * 解释: s2 包含 s1 的排列之一 ("ba").
 *
 *
 *
 *
 * 示例2:
 *
 *
 * 输入: s1= "ab" s2 = "eidboaoo"
 * 输出: False
 *
 *
 *
 *
 * 注意：
 *
 *
 * 输入的字符串只包含小写字母
 * 两个字符串的长度都在 [1, 10,000] 之间
 *
 *
 */

// @lc code=start

func checkInclusion(s1 string, s2 string) bool {
	// 通过直接比较哈希更有效率，因为输入的集合的k是有限个数（26个字母），这样每次比较最多比较26次
	if len(s1) > len(s2) {
		return false
	}
	h1 := make(map[string]int)
	h2 := make(map[string]int)
	for idx, c := range s1 {
		h1[string(c)]++
		h2[string(s2[idx])]++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		match := trueq
		for k, cnt := range h1 {
			if h2[k] != cnt {
				match = false
			}
		}

		if match {
			return true
		} else if (i + len(s1)) < len(s2) {
			h2[string(s2[i])]--
			h2[string(s2[i+len(s1)])]++
		}
	}
	return false
}

// @lc code=end