/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (35.13%)
 * Likes:    755
 * Dislikes: 0
 * Total Accepted:    143.9K
 * Total Submissions: 408.5K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	dislike := false
	prefix := strs[0]
	for i := 0; i < len(prefix) && !dislike; i++ {
		for _, str := range strs {
			if i > len(str)-1 || prefix[i] != str[i] {
				dislike = true
				break
			}
		}
		if !dislike {
			res += string(prefix[i])
		}
	}
	return res
}

// @lc code=end

