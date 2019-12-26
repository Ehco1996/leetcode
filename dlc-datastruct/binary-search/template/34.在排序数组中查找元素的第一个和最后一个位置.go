package medium

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (37.22%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    34.7K
 * Total Submissions: 92K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 如果数组中不存在目标值，返回 [-1, -1]。
 *
 * 示例 1:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 *
 * 示例 2:
 *
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 *
 */

func searchRange(nums []int, target int) []int {
	// 二分

	res := []int{}
	for l, r := 0, len(nums)-1; l <= r; {
		p := (l + r) / 2
		if nums[p] == target {
			startIdx := p
			endIdx := p
			for i := p; i < len(nums) && nums[i] == target; i++ {
				endIdx = i
			}
			for i := p; i >= 0 && nums[i] == target; i-- {
				startIdx = i
			}
			res = []int{startIdx, endIdx}
			break
		} else {
			if nums[p] < target {
				l = p + 1
			} else {
				r = p - 1
			}
		}
	}
	if len(res) > 0 {
		return res
	}
	return []int{-1, -1}

}
