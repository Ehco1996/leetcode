/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.30%)
 * Likes:    249
 * Dislikes: 0
 * Total Accepted:    43K
 * Total Submissions: 103.3K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */
import "sort"
import "math"

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func threeSumClosest(nums []int, target int) int {
	// 滑动窗口，从两边向中间收
	sort.Ints(nums)
	res := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if abs(sum, target) < abs(res, target) {
				res = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
