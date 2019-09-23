/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (46.39%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    14K
 * Total Submissions: 29.8K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 *
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * 给定 n 和 k，返回第 k 个排列。
 *
 * 说明：
 *
 *
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 *
 *
 * 示例 1:
 *
 * 输入: n = 3, k = 3
 * 输出: "213"
 *
 *
 * 示例 2:
 *
 * 输入: n = 4, k = 9
 * 输出: "2314"
 *
 *
 */

import "strconv"
import "strings"

func permute(nums []string) []string {
	res := []string{}

	if len(nums) == 0 {
		return res
	} else if len(nums) == 1 {
		return nums
	}

	for idx := 0; idx < len(nums); idx++ {
		last := []string{}
		last = append(last, nums[:idx]...)
		last = append(last, nums[idx+1:]...)
		for _, ans := range permute(last) {
			t := []string{nums[idx]}
			t = append(t, ans)
			res = append(res, strings.Join(t, ""))
		}
	}
	return res
}

func factorial(n int) int {
	res := 1
	for ; n > 0; n-- {
		res *= n
	}
	return res
}

func getPermutation(n int, k int) string {
	nums := make([]string, n)
	for i := 0; i < n; i++ {
		nums[i] = strconv.Itoa(i + 1)
	}
	if len(nums) == 1 {
		return nums[0]
	}

	cost := factorial(n - 1)
	ps := []string{}

	for idx, num := range nums {
		if k <= cost {
			last := []string{}
			last = append(last, nums[:idx]...)
			last = append(last, nums[idx+1:]...)
			ps = permute(last)
			for idx, v := range ps {
				ps[idx] = num + v
			}
			break
		}
		k -= cost

	}
	// fmt.Println(ps, cost)
	return ps[k-1]
}

