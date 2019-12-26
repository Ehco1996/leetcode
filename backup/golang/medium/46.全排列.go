package medium

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (70.87%)
 * Likes:    379
 * Dislikes: 0
 * Total Accepted:    46.5K
 * Total Submissions: 64.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */
func permute(nums []int) [][]int {
	// 迭代的方式
	res := [][]int{}
	if len(nums) == 0 {
		return res
	} else if len(nums) == 1 {
		res = append(res, nums)
		return res
	}

	for idx := 0; idx < len(nums); idx++ {
		last := []int{}
		last = append(last, nums[:idx]...)
		last = append(last, nums[idx+1:]...)
		for _, ans := range permute(last) {
			t := []int{nums[idx]}
			t = append(t, ans...)
			res = append(res, t)
		}
	}
	return res

}
