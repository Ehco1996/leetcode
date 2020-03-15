/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 *
 * https://leetcode-cn.com/problems/reverse-string/description/
 *
 * algorithms
 * Easy (67.22%)
 * Likes:    165
 * Dislikes: 0
 * Total Accepted:    71.5K
 * Total Submissions: 105.9K
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
 *
 * 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
 *
 * 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
 *
 *
 *
 * 示例 1：
 *
 * 输入：["h","e","l","l","o"]
 * 输出：["o","l","l","e","h"]
 *
 *
 * 示例 2：
 *
 * 输入：["H","a","n","n","a","h"]
 * 输出：["h","a","n","n","a","H"]
 *
 */

func helper(s []byte) []byte {
	l := len(s)
	if l <= 1 {
		return s
	} else {
		last := s[l-1]
		s = s[:l-1]
		s = helper(s)
		s = append([]byte{last}, s...)
		return s
	}
}

func reverseString(s []byte) {
	// 想要用递归做貌似不可能不用额外空间
	if len(s) <= 1 {
		return
	}

	for i, v := range helper(s) {
		s[i] = v
	}
}
