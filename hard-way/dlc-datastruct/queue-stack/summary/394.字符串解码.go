/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (46.86%)
 * Likes:    156
 * Dislikes: 0
 * Total Accepted:    11.5K
 * Total Submissions: 24.6K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 *
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 *
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 *
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 *
 * 示例:
 *
 *
 * s = "3[a]2[bc]", 返回 "aaabcbc".
 * s = "3[a2[c]]", 返回 "accaccacc".
 * s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".
 *
 *
 */

// @lc code=start
func decodeString(s string) string {
	r, _ := helper(s, 0)
	return r
}

func helper(s string, start int) (cur string, end int) {
	if len(s) == 0 {
		return "", 0
	}
	k := 0
	idx := start
	for idx < len(s) {
		ch := string(s[idx])
		if n, err := strconv.Atoi(ch); err == nil {
			k = k*10 + n
		} else if ch == "[" {
			decoded, end := helper(s, idx+1)
			cur += strings.Repeat(decoded, k)
			idx = end
			k = 0
		} else if ch == "]" {
			// this is the base case
			return cur, idx
		} else {
			cur += ch
		}
		idx++
	}
	return cur, idx
}

// @lc code=end

