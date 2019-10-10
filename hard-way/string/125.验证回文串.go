/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (40.54%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    50K
 * Total Submissions: 122.4K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 验证回文串示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */

// @lc code=start
import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// 双指针 主要去掉不是数字和字符的
	// 注意golang的string byte rune 之间的转换
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i <= j; {
		if !unicode.IsLetter(runes[i]) && !unicode.IsNumber(runes[i]) {
			i++
			continue
		}

		if !unicode.IsLetter(runes[j]) && !unicode.IsNumber(runes[j]) {
			j--
			continue
		}

		if runes[i] != runes[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

// @lc code=end

