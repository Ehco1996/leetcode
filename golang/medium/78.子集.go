/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (74.30%)
 * Likes:    349
 * Dislikes: 0
 * Total Accepted:    37.7K
 * Total Submissions: 50.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 *
 */

func subsets(nums []int) [][]int {
	res := [][]int{}
	size := len(nums)
	if size == 0 {
		return res
	}
	res = append(res, []int{})
	for _, num := range nums {
		for _, ans := range res {
			t := []int{num}
			t = append(t, ans...)
			res = append(res, t)
		}
	}
	return res
}

