/*
 * @lc app=leetcode.cn id=779 lang=golang
 *
 * [779] 第K个语法符号
 *
 * https://leetcode-cn.com/problems/k-th-symbol-in-grammar/description/
 *
 * algorithms
 * Medium (41.95%)
 * Likes:    55
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 17.8K
 * Testcase Example:  '1\n1'
 *
 * 在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。
 *
 * 给定行数 N 和序数 K，返回第 N 行中第 K个字符。（K从1开始）
 *
 *
 * 例子:
 *
 * 输入: N = 1, K = 1
 * 输出: 0
 *
 * 输入: N = 2, K = 1
 * 输出: 0
 *
 * 输入: N = 2, K = 2
 * 输出: 1
 *
 * 输入: N = 4, K = 5
 * 输出: 1
 *
 * 解释:
 * 第一行: 0
 * 第二行: 01
 * 第三行: 0110
 * 第四行: 01101001
 *
 *
 *
 * 注意：
 *
 *
 * N 的范围 [1, 30].
 * K 的范围 [1, 2^(N-1)].
 *
 *
 */

// @lc code=start
import "strconv"

func kthGrammar(N int, K int) int {

	cache := make([]string, 30)

	var helper func(n int) string
	helper = func(n int) string {
		if n == 1 {
			cache[n-1] = "0"
			return cache[n-1]
		}
		if cache[n-2] == "" {
			cache[n-2] = helper(n - 1)
		}
		now := ""
		for _, c := range cache[n-2] {
			s := string(c)
			if s == "0" {
				now += "01"
			} else if s == "1" {
				now += "10"
			} else {
				now += s
			}
		}
		cache[n-1] = now
		return now
	}

	now := helper(N)
	res, _ := strconv.Atoi(string(now[K-1]))
	return res
}

// @lc code=end

