/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 *
 * https://leetcode-cn.com/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (37.81%)
 * Likes:    71
 * Dislikes: 0
 * Total Accepted:    6.6K
 * Total Submissions: 17.6K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典
 * wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
 *
 * 说明：
 *
 *
 * 分隔时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入:
 * s = "catsanddog"
 * wordDict = ["cat", "cats", "and", "sand", "dog"]
 * 输出:
 * [
 * "cats and dog",
 * "cat sand dog"
 * ]
 *
 *
 * 示例 2：
 *
 * 输入:
 * s = "pineapplepenapple"
 * wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
 * 输出:
 * [
 * "pine apple pen apple",
 * "pineapple pen apple",
 * "pine applepen apple"
 * ]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入:
 * s = "catsandog"
 * wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出:
 * []
 *
 *
 */

// @lc code=start
import "strings"

func wordBreak(s string, wordDict []string) []string {
	// dp dp[i] 记录 s[:i] 是否满足条件
	// dfs 拿出所有结果
	h := make(map[string]bool)
	for _, sub := range wordDict {
		h[sub] = true
	}

	dp := make([]bool, len(s)+1)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			if i > 0 && !dp[i] {
				// 表明之s[i] 已经不能构成了
				break
			}
			if h[s[i:j]] {
				// s[i:j] 在worddict里
				dp[j] = true
			}
		}
	}

	res := []string{}
	if !dp[len(s)] {
		return res
	}
	var dfs = func(string, []string) {}
	dfs = func(last string, path []string) {
		if len(last) == 0 {
			res = append(res, strings.Join(path, " "))
			return
		}
		for i := 1; i <= len(last); i++ {
			if h[last[:i]] {
				dfs(last[i:], append(path, last[:i]))
			}
		}
	}
	dfs(s, []string{})
	return res
}

// @lc code=end

