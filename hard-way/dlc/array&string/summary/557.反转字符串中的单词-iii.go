/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 *
 * https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (68.18%)
 * Likes:    125
 * Dislikes: 0
 * Total Accepted:    32.9K
 * Total Submissions: 48.2K
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

// @lc code=start
func reverse(nums []byte) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
func reverseWords(s string) string {
	bs := []byte(s)
	last := 0
	for i := 0; i < len(bs); i++ {
		if string(bs[i]) == " " {
			reverse(bs[last:i])
			last = i + 1
		} else if i == len(bs)-1 {
			reverse(bs[last:])
		}
	}
	return string(bs)
}

// @lc code=end

