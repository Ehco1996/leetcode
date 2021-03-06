/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 *
 * https://leetcode-cn.com/problems/russian-doll-envelopes/description/
 *
 * algorithms
 * Hard (32.71%)
 * Likes:    96
 * Dislikes: 0
 * Total Accepted:    6.3K
 * Total Submissions: 18.9K
 * Testcase Example:  '[[5,4],[6,4],[6,7],[2,3]]'
 *
 * 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h)
 * 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
 *
 * 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
 *
 * 说明:
 * 不允许旋转信封。
 *
 * 示例:
 *
 * 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
 * 输出: 3
 * 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 *
 *
 */

// @lc code=start
import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxEnvelopes(envelopes [][]int) int {
	// dp dp(j) 表示当前信封能包含的最大的信封数量
	// 先将输入数组按wh排序，保证能包含当前信封的信封一定在当前信封之后，
	// 然后从前往后更新dp，如果遇到能包含当前信封i的信封j，
	// 则dp[j]=max(dp[j],dp[i]+1)

	if len(envelopes) == 0 {
		return 0
	}

	// sort by interval[0]
	sort.SliceStable(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		for j := i; j < len(envelopes); j++ {
			if envelopes[j][0] > envelopes[i][0] && envelopes[j][1] > envelopes[i][1] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	ans := dp[0]
	for _, v := range dp {
		ans = max(v, ans)
	}
	// +1 是因为自己也是一个信封
	return ans + 1
}

// @lc code=end

