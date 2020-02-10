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

func n2Array(n int) []string {
	res := []string{}

	for i := 1; i <= n; i++ {
		res = append(res, strconv.Itoa(i))
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
	// 不需要生成所有的排列
	// n 个数字有 n！种全排列，每种数字开头的全排列有 (n - 1)!种。
	// 所以用 k / (n - 1)! 就可以得到第 k 个全排列是以第几个数字开头的。
	// 用 k % (n - 1)! 就可以得到第 k 个全排列是某个数字开头的全排列中的第几个

	nums := n2Array(n)
	if k == 1 {
		return strings.Join(nums, "")
	}
	fact := factorial(n - 1)

	round := n - 1
	k -= 1
	res := []string{}
	for round >= 0 {
		idx := k / fact
		k = k % fact
		res = append(res, nums[idx])
		// Remove the element at index idx from nums
		copy(nums[idx:], nums[idx+1:]) // Shift nums[idx+1:] left one index.
		nums[len(nums)-1] = ""         // Erase last element (write zero value).
		nums = nums[:len(nums)-1]      // Truncate slice.
		if round > 0 {
			fact /= round
		}
		round -= 1
	}
	return strings.Join(res, "")
}
