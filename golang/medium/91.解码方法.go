/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 *
 * https://leetcode-cn.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (21.44%)
 * Likes:    196
 * Dislikes: 0
 * Total Accepted:    17.1K
 * Total Submissions: 78.1K
 * Testcase Example:  '"12"'
 *
 * 一条包含字母 A-Z 的消息通过以下方式进行了编码：
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * 给定一个只包含数字的非空字符串，请计算解码方法的总数。
 *
 * 示例 1:
 *
 * 输入: "12"
 * 输出: 2
 * 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
 *
 *
 * 示例 2:
 *
 * 输入: "226"
 * 输出: 3
 * 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
 *
 *
 */
import "strconv"

func numDecodings(s string) int {
	// d(si) = dp(s[:i-2]) + dp(s[:i-1])
	if len(s) == 0 || string(s[0]) == "0" {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		// 如果该位不为“0” 则合法
		if string(s[i-1]) != "0" {
			dp[i] += dp[i-1]
		}
		//如果后两位能组成"1x"（x为任意数字）或者"2x"（x小于7），说明最后两位组成字母合法
		num, _ := strconv.Atoi(string(s[i-1]))
		if string(s[i-2]) == "1" || (string(s[i-2]) == "2" && num <= 6) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

