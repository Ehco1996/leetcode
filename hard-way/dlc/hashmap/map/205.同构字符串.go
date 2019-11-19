/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (46.26%)
 * Likes:    143
 * Dislikes: 0
 * Total Accepted:    20.9K
 * Total Submissions: 45K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
 *
 * 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
 *
 * 示例 1:
 *
 * 输入: s = "egg", t = "add"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "foo", t = "bar"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: s = "paper", t = "title"
 * 输出: true
 *
 * 说明:
 * 你可以假设 s 和 t 具有相同的长度。
 *
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	// hash k: 字符 v:上一次出现的idx
	// 每次比较出现的 idx是否一样
	h1 := make(map[byte]int)
	h2 := make(map[byte]int)

	for idx := 0; idx < len(s); idx++ {
		c1 := s[idx]
		c2 := t[idx]
		v1, ok1 := h1[c1]
		v2, ok2 := h2[c2]
		if v1 != v2 || (ok1 != ok2) {
			return false
		} else {
			h1[c1] = idx
			h2[c2] = idx
		}

	}
	return true
}

// @lc code=end

