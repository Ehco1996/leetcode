package medium

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (31.30%)
 * Likes:    275
 * Dislikes: 0
 * Total Accepted:    24.2K
 * Total Submissions: 75.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 *
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 *
 * 必须原地修改，只允许使用额外常数空间。
 *
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */
import "sort"

func nextPermutation(nums []int) {
	// 从后向前找，找到第一个比最后一个数要小的数，交换两数的位置，然后将这个数往后的所有数升序排列
	if len(nums) <= 1 {
		return
	}

	target := nums[len(nums)-1]
	for idx := len(nums) - 2; idx >= 0; idx-- {
		if nums[idx] < target {
			// 排序后找到第一个比当前数大的数，交换位置
			sort.Ints(nums[idx+1:])
			for i := idx + 1; i < len(nums); i++ {
				if nums[i] > nums[idx] {
					nums[i], nums[idx] = nums[idx], nums[i]
					return
				}
			}
		} else {
			target = nums[idx]
		}
	}
	sort.Ints(nums)
}
