/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (34.10%)
 * Likes:    247
 * Dislikes: 0
 * Total Accepted:    24.8K
 * Total Submissions: 69.2K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 *
 * 示例 1:
 *
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3
 * 解释: 11 = 5 + 5 + 1
 *
 * 示例 2:
 *
 * 输入: coins = [2], amount = 3
 * 输出: -1
 *
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 *
 */

// @lc code=start
import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	// 拆分子问题时，会发现组合的多少之和当前币种的选择有关
	// f(n)=min{f(n-1),f(n-5),f(n-11)}+1

	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		cost := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				if dp[i-coins[j]] != math.MaxInt32 {
					cost = min(cost, dp[i-coins[j]]+1)
				}
			}
		}
		dp[i] = cost
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// @lc code=end

