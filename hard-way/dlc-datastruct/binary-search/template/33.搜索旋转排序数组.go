/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.22%)
 * Likes:    442
 * Dislikes: 0
 * Total Accepted:    59.1K
 * Total Submissions: 163.2K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 *
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 *
 * 你可以假设数组中不存在重复的元素。
 *
 * 你的算法时间复杂度必须是 O(log n) 级别。
 *
 * 示例 1:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	// 两次二分查找，一次二分找到旋转点，另外一次找到目标值
	var bs func(l, r int) int
	bs = func(l, r int) int {
		for l <= r {
			p := (l + r) / 2
			if nums[p] == target {
				return p
			} else {
				if nums[p] < target {
					l = p + 1
				} else {
					r = p - 1
				}
			}
		}
		return -1
	}

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// find rotate_index
	rotateIndex := 0
	if nums[0] > nums[len(nums)-1] {
		for l, r := 0, len(nums)-1; l <= r; {
			p := (l + r) / 2
			if nums[p] > nums[p+1] {
				// 从升序变成降序了 找到了
				rotateIndex = p + 1
				break
			} else {
				if nums[p] < nums[l] {
					r = p - 1
				} else {
					l = p + 1
				}
			}
		}
	}

	if nums[rotateIndex] == target {
		return rotateIndex
	}
	if rotateIndex == 0 {
		return bs(0, len(nums)-1)
	}
	if target < nums[0] {
		return bs(rotateIndex, len(nums)-1)
	}
	return bs(0, rotateIndex)
}

// @lc code=end

