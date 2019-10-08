/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (50.76%)
 * Likes:    209
 * Dislikes: 0
 * Total Accepted:    21.8K
 * Total Submissions: 42.9K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 *
 * 示例 1:
 *
 * 输入: n = 12
 * 输出: 3
 * 解释: 12 = 4 + 4 + 4.
 *
 * 示例 2:
 *
 * 输入: n = 13
 * 输出: 2
 * 解释: 13 = 4 + 9.
 *
 */

// @lc code=start
import "sort"

func numSquares1(n int) int {
	// 暴力 超时
	if n == 0 || n == 1 {
		return n
	}

	h := make(map[int]bool)
	for i := 1; i*i <= n; i++ {
		h[i*i] = true
	}

	ans := []int{}
	for num, _ := range h {
		last := n - num
		if h[last] {
			ans = append(ans, 1+numSquares(n-num))
		}
	}
	sort.Ints(ans)
	return ans[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func numSquares(n int) int {
	// dp[i] = MIN(dp[i], dp[i - j * j] + 1)
	// i表示当前数字，j*j表示平方数(从1到最大的平方数)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 最坏的情况
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]

}

// @lc code=end

