package medium

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (33.93%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    10.6K
 * Total Submissions: 31.1K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 *
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 *
 * 示例 1:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 *
 * 进阶:
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
 */
func search(nums []int, target int) bool {
	var bs func(l, r int) bool
	bs = func(l, r int) bool {
		for l <= r {
			p := (l + r) / 2
			if nums[p] == target {
				return true
			} else {
				if nums[p] < target {
					l = p + 1
				} else {
					r = p - 1
				}
			}
		}
		return false
	}

	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}

	// find rotate_index O（N）
	rotateIndex := 1
	for rotateIndex < len(nums) {
		if nums[rotateIndex] < nums[rotateIndex-1] {
			break
		}
		rotateIndex++
	}
	if rotateIndex >= len(nums) {
		rotateIndex = 0
	}

	if nums[rotateIndex] == target {
		return true
	}
	if rotateIndex == 0 {
		return bs(0, len(nums)-1)
	}
	if target < nums[0] {
		return bs(rotateIndex, len(nums)-1)
	}
	return bs(0, rotateIndex)

}
