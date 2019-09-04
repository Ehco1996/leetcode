package main

import "fmt"

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 *
 * https://leetcode-cn.com/problems/reverse-vowels-of-a-string/description/
 *
 * algorithms
 * Easy (47.68%)
 * Likes:    56
 * Dislikes: 0
 * Total Accepted:    14.3K
 * Total Submissions: 29.9K
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

func IsVowel(s byte) bool {
	al := []byte("aAeEiIoOuU")
	for _, c := range al {
		if c == s {
			return true
		}
	}
	return false
}

func reverseVowels(s string) string {
	bs := []byte(s)

	for i, j := 0, len(bs)-1; i < j; {

		if IsVowel(bs[i]) &&IsVowel(bs[j]) {
			bs[i], bs[j] = bs[j], bs[i]
			i++
			j--
		}
		if !IsVowel(bs[i]) {
			i++
		}
		if !IsVowel(bs[j]) {
			j--
		}
	}
	return string(bs)
}

// func main() {
// 	fmt.Println(reverseVowels("aA"))
// }
