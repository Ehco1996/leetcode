package easy

import "strings"

/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 *
 * https://leetcode-cn.com/problems/keyboard-row/description/
 *
 * algorithms
 * Easy (66.95%)
 * Likes:    54
 * Dislikes: 0
 * Total Accepted:    9.6K
 * Total Submissions: 14.4K
 * Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
 *
 * 给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
 *
 *
 *
 *
 *
 *
 *
 * 示例：
 *
 * 输入: ["Hello", "Alaska", "Dad", "Peace"]
 * 输出: ["Alaska", "Dad"]
 *
 *
 *
 *
 * 注意：
 *
 *
 * 你可以重复使用键盘上同一字符。
 * 你可以假设输入的字符串将只包含字母。
 *
 */

func findWords(words []string) []string {
	// 哈希
	res := []string{}

	h := map[string]int{
		"q": 1,
		"w": 1,
		"e": 1,
		"r": 1,
		"t": 1,
		"y": 1,
		"u": 1,
		"i": 1,
		"o": 1,
		"p": 1,

		"a": 2,
		"s": 2,
		"d": 2,
		"f": 2,
		"g": 2,
		"h": 2,
		"j": 2,
		"k": 2,
		"l": 2,

		"z": 3,
		"x": 3,
		"c": 3,
		"v": 3,
		"b": 3,
		"n": 3,
		"m": 3,
	}

	for _, word := range words {
		line := h[strings.ToLower(string(word[0]))]
		pass := true
		for _, c := range word[1:] {
			if h[strings.ToLower(string(c))] != line {
				pass = false
				break
			}
		}
		if pass {
			res = append(res, word)
		}
	}
	return res
}
