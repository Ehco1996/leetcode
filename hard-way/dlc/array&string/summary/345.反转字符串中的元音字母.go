/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (48.21%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    17.9K
 * Total Submissions: 37.2K
 * Testcase Example:  '"hello"'
 *
 * 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
 *
 * 示例 1:
 *
 * 输入: "hello"
 * 输出: "holle"
 *
 *
 * 示例 2:
 *
 * 输入: "leetcode"
 * 输出: "leotcede"
 *
 * 说明:
 * 元音字母不包含字母"y"。
 *
 */

// @lc code=start

func isVowel(in byte) bool {
	s := string(in)
	if s != "a" && s != "e" && s != "i" && s != "o" && s != "u" && s != "A" && s != "E" && s != "I" && s != "O" && s != "U" {
		return false
	}
	return true
}

func reverseVowels(s string) string {

	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; {
		if isVowel(bs[i]) && isVowel(bs[j]) {
			bs[i], bs[j] = bs[j], bs[i]
			i++
			j--
			continue
		}
		if !isVowel(bs[i]) {
			i++
		}
		if !isVowel(bs[j]) {
			j--
		}
	}
	return string(bs)
}

// @lc code=end

