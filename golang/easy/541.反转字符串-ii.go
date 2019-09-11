// package easy

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 *
 * https://leetcode-cn.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (47.95%)
 * Likes:    43
 * Dislikes: 0
 * Total Accepted:    6.8K
 * Total Submissions: 14K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k
 * 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。
 *
 * 示例:
 *
 *
 * 输入: s = "abcdefg", k = 2
 * 输出: "bacdfeg"
 *
 *
 * 要求:
 *
 *
 * 该字符串只包含小写的英文字母。
 * 给定字符串的长度和 k 在[1, 10000]范围内。
 *
 *
 */

func reverse(b []byte) {
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left, right = left+1, right-1
	}
}
func reverseStr(s string, k int) string {
	sb := []byte(s)
	for i := 0; i < len(sb); i = i + 2*k {
		if i+k < len(sb) {
			reverse(sb[i : i+k])
		} else {
			reverse(sb[i:])
		}
	}
	return string(sb)
}
