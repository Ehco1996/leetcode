/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (55.84%)
 * Likes:    182
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 36.9K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 根据每日 气温 列表，请重新生成一个列表，对应位置的输入是你需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
 *
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4,
 * 2, 1, 1, 0, 0]。
 *
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
 *
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	//递减的stack
	res := make([]int, len(T))
	stack := []int{}
	for idx, t := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < t {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[cur] = idx - cur
		}
		stack = append(stack, idx)
	}
	return res
}

// @lc code=end

