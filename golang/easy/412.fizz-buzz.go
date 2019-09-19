package easy

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode-cn.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (60.74%)
 * Likes:    33
 * Dislikes: 0
 * Total Accepted:    18.9K
 * Total Submissions: 31K
 * Testcase Example:  '1'
 *
 * 写一个程序，输出从 1 到 n 数字的字符串表示。
 *
 * 1. 如果 n 是3的倍数，输出“Fizz”；
 *
 * 2. 如果 n 是5的倍数，输出“Buzz”；
 *
 * 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 *
 * 示例：
 *
 * n = 15,
 *
 * 返回:
 * [
 * ⁠   "1",
 * ⁠   "2",
 * ⁠   "Fizz",
 * ⁠   "4",
 * ⁠   "Buzz",
 * ⁠   "Fizz",
 * ⁠   "7",
 * ⁠   "8",
 * ⁠   "Fizz",
 * ⁠   "Buzz",
 * ⁠   "11",
 * ⁠   "Fizz",
 * ⁠   "13",
 * ⁠   "14",
 * ⁠   "FizzBuzz"
 * ]
 *
 *
 */
func fizzBuzz(n int) []string {
	// ???
	res := []string{}
	for num := 1; num <= n; num++ {
		if num%3 == 0 && num%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if num%3 == 0 {
			res = append(res, "Fizz")
		} else if num%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(num))
		}
	}
	return res
}
