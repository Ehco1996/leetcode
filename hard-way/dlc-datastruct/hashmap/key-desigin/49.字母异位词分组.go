/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (58.99%)
 * Likes:    210
 * Dislikes: 0
 * Total Accepted:    36.5K
 * Total Submissions: 61.5K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */

// @lc code=start
import "sort"

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	// 当且仅当它们的排序字符串相等时，两个字符串是字母异位词。
	res := [][]string{}
	if len(strs) == 0 {
		return res
	}

	h := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		h[key] = append(h[key], str)
	}

	for _, v := range h {
		res = append(res, v)
	}
}

// @lc code=end

