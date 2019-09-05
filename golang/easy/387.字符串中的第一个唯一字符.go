// package golang

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (39.70%)
 * Likes:    140
 * Dislikes: 0
 * Total Accepted:    40.5K
 * Total Submissions: 100.6K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 * 案例:
 *
 *
 * s = "leetcode"
 * 返回 0.
 *
 * s = "loveleetcode",
 * 返回 2.
 *
 *
 *
 *
 * 注意事项：您可以假定该字符串只包含小写字母。
 *
 */
func firstUniqChar(s string) int {
	h := make(map[rune][]int)
	for idx,c := range s{
		h[c] = append(h[c],idx)
	}

	minIdx := len(s) +1

	for _,v :=range h{
		if len(v)==1&&v[0]<minIdx{
			minIdx = v[0]
		}
	}

	if minIdx == len(s)+1{
		return -1
	}else{
		return minIdx
	}



}
