package easy

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (66.79%)
 * Likes:    119
 * Dislikes: 0
 * Total Accepted:    27.4K
 * Total Submissions: 40.6K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
 *
 * 示例 1:
 *
 *
 * 输入: "Let's take LeetCode contest"
 * 输出: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 * 注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
 *
 */

func reverse(b []byte) {
	left, right := 0, len(b)-1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left, right = left+1, right-1
	}
}

func reverseWords(s string) string {
	// 双指针
	sb := []byte(s)
	last := 0
	for i := 0; i < len(sb); i++ {
		if string(sb[i]) == " " {
			reverse(sb[last:i])
			last = i + 1
		} else if i == len(sb)-1 {
			reverse(sb[last:])
		}
	}
	return string(sb)
}
