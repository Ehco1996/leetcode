package easy

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (54.67%)
 * Likes:    106
 * Dislikes: 0
 * Total Accepted:    44.1K
 * Total Submissions: 79.6K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */
func isAnagram(s string, t string) bool {
	//hash
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, c := range s {
		m[c] += 1
	}

	for _, c := range t {
		v, ok := m[c]
		if !ok || v < 1 {
			return false
		} else {
			m[c] -= 1
		}
	}
	return true

}
