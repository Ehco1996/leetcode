/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 *
 * https://leetcode-cn.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (33.22%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    5.1K
 * Total Submissions: 15K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * 给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
 *
 * 示例 1:
 *
 * 输入: nums = [1, 5, 1, 1, 6, 4]
 * 输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
 *
 * 示例 2:
 *
 * 输入: nums = [1, 3, 2, 2, 3, 1]
 * 输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
 *
 * 说明:
 * 你可以假设所有输入都会得到有效的结果。
 *
 * 进阶:
 * 你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
 *
 */

// @lc code=start
import "sort"

func wiggleSort(nums []int) {
	// 数组排序后分为两个部分。然后一个一个往原来的数组里插入

	bak := make([]int, len(nums))
	copy(bak, nums)
	sort.Ints(bak)
	l := 0
	r := len(nums) - 1
	if len(nums)%2 == 0 {
		l = len(nums)/2 - 1
	} else {
		l = len(nums) / 2
	}

	useLeft := true
	for i := 0; i < len(nums); i++ {
		if useLeft {
			nums[i] = bak[l]
			l--
			useLeft = false
		} else {
			nums[i] = bak[r]
			r--
			useLeft = true
		}
	}
}

// @lc code=end

