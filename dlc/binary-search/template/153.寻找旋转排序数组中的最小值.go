/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (49.94%)
 * Likes:    112
 * Dislikes: 0
 * Total Accepted:    24.6K
 * Total Submissions: 49.2K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 你可以假设数组中不存在重复元素。
 *
 * 示例 1:
 *
 * 输入: [3,4,5,1,2]
 * 输出: 1
 *
 * 示例 2:
 *
 * 输入: [4,5,6,7,0,1,2]
 * 输出: 0
 *
 */

// @lc code=start
func findMin(nums []int) int {
	// 二分 旋转点就是最小值
	left, right := 0, len(nums)-1
	var mid int
	var min int
	if nums[left] > nums[right] {
		for left <= right {
			mid = (left + right) / 2
			if nums[mid] > nums[mid+1] {
				min = nums[mid+1]
				break
			}
			if nums[mid] < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	} else {
		min = nums[0]
	}
	return min
}

// @lc code=end

