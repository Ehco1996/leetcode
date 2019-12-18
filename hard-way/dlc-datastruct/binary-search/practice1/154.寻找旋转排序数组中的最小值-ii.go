/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Hard (45.85%)
 * Likes:    57
 * Dislikes: 0
 * Total Accepted:    10.4K
 * Total Submissions: 22.5K
 * Testcase Example:  '[1,3,5]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] )。
 *
 * 请找出其中最小的元素。
 *
 * 注意数组中可能存在重复的元素。
 *
 * 示例 1：
 *
 * 输入: [1,3,5]
 * 输出: 1
 *
 * 示例 2：
 *
 * 输入: [2,2,2,0,1]
 * 输出: 0
 *
 * 说明：
 *
 *
 * 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
 * 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
 *
 *
 */

// @lc code=start
func findMin(nums []int) int {
	//1. nums[mid] < nums[right]，这种情况是当前的数小于右边数组的一个数，说明当前的位置肯定是在右半数组中，因此可以往左查找[left, mid]
	//2. nums[mid] > nums[right]，这种情况是当前的数大于右边数组的一个数，说明当前的位置在左半数组中，因此要查找右半数组第一个数的位置可以向右搜索[mid+1, right]
	//3. nums[mid] = nums[right]，这种情况下无法判断当前位置是在左边还是右边，但是让右指针左移一位并不会影响到最小值，因为nums[mid] = nums[right], 因此搜索[left, right-1]

	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[l]
}

// @lc code=end

