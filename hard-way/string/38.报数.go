/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 *
 * https://leetcode-cn.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (53.31%)
 * Likes:    407
 * Dislikes: 0
 * Total Accepted:    77.8K
 * Total Submissions: 142.7K
 * Testcase Example:  '1'
 *
 * 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 * 1 被读作  "one 1"  ("一个一") , 即 11。
 * 11 被读作 "two 1s" ("两个一"）, 即 21。
 * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
 *
 * 给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
 *
 * 注意：整数序列中的每一项将表示为一个字符串。
 *
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "1"
 * 解释：这是一个基本样例。
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "1211"
 * 解释：当 n = 3 时，序列是 "21"，其中我们有 "2" 和 "1" 两组，"2" 可以读作 "12"，也就是出现频次 = 1 而 值 =
 * 2；类似 "1" 可以读作 "11"。所以答案是 "12" 和 "11" 组合在一起，也就是 "1211"。
 *
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return genByLast(countAndSay(n - 1))
}

func genByLast(s string) string {
	// 数第一个字符出现了几次，然后拼上接下来的
	if s == "" {
		return ""
	}
	idx := 1
	for ; idx < len(s) && s[idx] == s[0]; idx++ {
	}
	pre := strconv.Itoa(idx)
	return pre + string(s[0]) + genByLast(string(s[idx:]))

}

// @lc code=end

