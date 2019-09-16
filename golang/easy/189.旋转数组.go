package easy

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Easy (38.24%)
 * Likes:    356
 * Dislikes: 0
 * Total Accepted:    63.9K
 * Total Submissions: 165.6K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5,6,7] 和 k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 * 输入: [-1,-100,3,99] 和 k = 2
 * 输出: [3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 * 说明:
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 要求使用空间复杂度为 O(1) 的 原地 算法。
 *
 *
 */
func rotate1(nums []int, k int) {
	// 暴力移动
	lens := len(nums)
	for c := 0; c < k; c++ {
		tmp := nums[lens-1]
		var now int
		for i := 0; i < lens; i++ {
			now = nums[i]
			nums[i] = tmp
			tmp = now
		}
	}
}

func reverse(nums []int, i, j int) {
	for i != j && i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	// 只需要将数组的部分翻转三次
	lens := len(nums)
	if k >= lens {
		k = k % lens
	}
	reverse(nums, 0, lens-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, lens-1)
}
