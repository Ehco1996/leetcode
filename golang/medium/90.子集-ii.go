package medium

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (56.22%)
 * Likes:    126
 * Dislikes: 0
 * Total Accepted:    13.6K
 * Total Submissions: 23.7K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 */

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// 回溯
	res := [][]int{}
	size := len(nums)
	if size == 0 {
		return res
	}
	sort.Ints(nums)

	path := []int{}
	helper(nums, 0, path, &res)
	return res

}

func helper(nums []int, start int, path []int, res *[][]int) {
	t := make([]int, len(path))
	copy(t, path)
	*res = append(*res, t)

	for i := start; i < len(nums); i++ {
		// 如果和上个数子相等就跳过（剪枝）
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		helper(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
