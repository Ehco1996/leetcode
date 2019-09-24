/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (69.76%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    20K
 * Total Submissions: 28.4K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 *
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 *
 */
func combine(n int, k int) [][]int {

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	var helper func(nums []int, k int) [][]int
	helper = func(nums []int, k int) [][]int {
		res := [][]int{}
		if k == 0 {
			return res
		} else if k == 1 {
			for _, num := range nums {
				res = append(res, []int{num})
			}
			return res
		}

		for idx, num := range nums {
			for _, ans := range helper(nums[idx+1:], k-1) {
				t := []int{num}
				t = append(t, ans...)
				res = append(res, t)
			}
		}
		return res
	}
	return helper(nums, k)
}

