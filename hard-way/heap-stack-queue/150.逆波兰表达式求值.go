/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 *
 * https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (45.21%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    16K
 * Total Submissions: 34.5K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * 根据逆波兰表示法，求表达式的值。
 *
 * 有效的运算符包括 +, -, *, / 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
 *
 * 说明：
 *
 *
 * 整数除法只保留整数部分。
 * 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
 *
 *
 * 示例 1：
 *
 * 输入: ["2", "1", "+", "3", "*"]
 * 输出: 9
 * 解释: ((2 + 1) * 3) = 9
 *
 *
 * 示例 2：
 *
 * 输入: ["4", "13", "5", "/", "+"]
 * 输出: 6
 * 解释: (4 + (13 / 5)) = 6
 *
 *
 * 示例 3：
 *
 * 输入: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * 输出: 22
 * 解释:
 * ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 *
 */

// @lc code=start
func evalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		switch ch := tokens[i]; ch {
		case "+", "-", "*", "/":
			int1, int2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var res int
			if ch == "+" {
				res = int1 + int2
			} else if ch == "-" {
				res = int1 - int2
			} else if ch == "*" {
				res = int1 * int2
			} else {
				res = int1 / int2
			}
			stack = append(stack, res)
		default:
			i, _ := strconv.Atoi(ch)
			stack = append(stack, i)
		}
	}
	return stack[0]
}

// @lc code=end

