/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.10%)
 * Likes:    357
 * Dislikes: 0
 * Total Accepted:    43.6K
 * Total Submissions: 120.8K
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

func bs(nums []int, target int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	// 两次二分查找，一次二分找到旋转点，另外一次找到目标值

	if len(nums) == 0 {
		return -1
	}

	var pIdx int
	// 发生了旋转才需要搜索
	for l, r := 0, len(nums)-1; l <= r && nums[0] > nums[len(nums)-1]; {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			pIdx = mid
			break
		}
		if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	res := bs(nums[:pIdx+1], target)
	if res == -1 {
		if t := bs(nums[pIdx+1:], target); t != -1 {
			res = t + pIdx + 1
		}
	}
	return res
}


