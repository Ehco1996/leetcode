/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (51.25%)
 * Likes:    243
 * Dislikes: 0
 * Total Accepted:    26.8K
 * Total Submissions: 52.3K
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
func numSquares(n int) int {
	// bfs 每层都找比现在更小的所有完全平方数 需要剪枝
	// 其实这也是一道dp题，解法可以搜之前的答案
	queue := []int{n}
	res := 0

	for len(queue) > 0 {
		// 这里剪枝，去掉每层出来的重复元素
		h := make(map[int]bool)

		for len(queue) > 0 {
			target := queue[0]
			queue = queue[1:]

			if target == 0 {
				// 如果当前队列里的数是0说明一件分解完成了
				return res
			}

			i := 1
			for i*i <= target {
				h[target-i*i] = true
				i++
			}
		}
		res++
		for k, _ := range h {
			queue = append(queue, k)
		}
	}
	return res
}

// @lc code=end

