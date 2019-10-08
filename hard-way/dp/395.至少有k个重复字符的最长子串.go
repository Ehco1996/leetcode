/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/description/
 *
 * algorithms
 * Medium (40.19%)
 * Likes:    68
 * Dislikes: 0
 * Total Accepted:    4.2K
 * Total Submissions: 10.3K
 * Testcase Example:  '"aaabb"\n3'
 *
 * 找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。
 *
 * 示例 1:
 *
 *
 * 输入:
 * s = "aaabb", k = 3
 *
 * 输出:
 * 3
 *
 * 最长子串为 "aaa" ，其中 'a' 重复了 3 次。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s = "ababbc", k = 2
 *
 * 输出:
 * 5
 *
 * 最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
 *
 *
 */

// @lc code=start
import (
	"sort"
	"strings"
)

func longestSubstring(s string, k int) int {
	// 统计每个字符出现的次数，如果出现的次数小于k 则包含该字符的子串不符合，将子串分割为左右两个部分

	cntMap := make(map[rune]int)
	for _, c := range s {
		cntMap[c]++
	}
	for key, cnt := range cntMap {
		if cnt < k {
			temp := []int{}
			for _, p := range strings.Split(s, string(key)) {
				temp = append(temp, longestSubstring(p, k))
			}

			sort.Ints(temp)
			return temp[len(temp)-1]
		}
	}
	return len(s)
}

// @lc code=end

