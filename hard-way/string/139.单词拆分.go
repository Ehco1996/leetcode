/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (42.28%)
 * Likes:    206
 * Dislikes: 0
 * Total Accepted:    20.4K
 * Total Submissions: 47.7K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 *
 * 说明：
 *
 *
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// dp dp[i] 记录 s[:i] 是否满足条件
	h := make(map[string]int)
	for _, sub := range wordDict {
		h[sub]++
	}

	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if i > 0 && dp[i] != 1 {
				// 表明s[i]已经不能构成了
				break
			}
			if _, ok := h[s[i:j]]; ok {
				// s[i:j] 在worddict里
				dp[j] = 1
			}
		}
	}
	if dp[len(s)] == 1 {
		return true
	}
	return false
}

// @lc code=end

