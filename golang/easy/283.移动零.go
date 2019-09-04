package main

import "fmt"

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (55.76%)
 * Likes:    381
 * Dislikes: 0
 * Total Accepted:    65.3K
 * Total Submissions: 115.7K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */
func moveZeroes(nums []int) {
	// 找到所有不是零的数安顺序排好，末尾补零就可以
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}
	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}

// func main() {
// 	l := []int{0, 1, 0, 3, 12}
// 	moveZeroes(l)
// 	fmt.Println(l)
// }
