// package golang

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (31.53%)
 * Likes:    38
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 23.9K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 *
 * 示例:
 *
 * 输入: "Hello, my name is John"
 * 输出: 5
 *
 *
 */
 func countSegments(s string) int {

	count := 1
	found := false
	for i := 0; i < len(s); i++ {
		if i > 0 && string(s[i]) == " " && string(s[i-1]) != " " {
			count++
		}
		if string(s[i]) != " " {
			found = true
		}
		if i == len(s)-1 && string(s[i]) == " " {
			count--
		}
	}
	if found {
		return count
	}
	return 0
}

