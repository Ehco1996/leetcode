/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (53.40%)
 * Likes:    159
 * Dislikes: 0
 * Total Accepted:    22.7K
 * Total Submissions: 41.9K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 *
 * 示例:
 *
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 *
 */
import "sort"

func permuteUnique(nums []int) [][]int {
	// 迭代的方式 注意去重
	res := [][]int{}
	if len(nums) == 0 {
		return res
	} else if len(nums) == 1 {
		res = append(res, nums)
		return res
	}

	sort.Ints(nums)

	for idx := 0; idx < len(nums); idx++ {
		last := []int{}
		last = append(last, nums[:idx]...)
		last = append(last, nums[idx+1:]...)
		for _, ans := range permuteUnique(last) {
			if idx > 0 && nums[idx] == nums[idx-1] {
				continue
			}
			t := []int{nums[idx]}
			t = append(t, ans...)
			res = append(res, t)
		}
	}
	return res
}

