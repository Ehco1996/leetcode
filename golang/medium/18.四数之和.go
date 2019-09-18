// package medium

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (35.94%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    31.7K
 * Total Submissions: 88.2K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for l, r := j+1, len(nums)-1; l < r; {
				temp := nums[i] + nums[j] + nums[l] + nums[r]
				if temp == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
				} else if temp > target {
					r--
				} else {
					l++
				}
			}
		}
	}

	// repeated
	h := make(map[string]int)
	clean := [][]int{}
	for _, ans := range res {
		key := ""
		for _, num := range ans {
			key += strconv.Itoa(num)
		}
		if _, ok := h[key]; !ok {
			h[key]++
			clean = append(clean, ans)
		}
	}
	return clean
}
