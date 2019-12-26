/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (45.61%)
 * Likes:    342
 * Dislikes: 0
 * Total Accepted:    83.3K
 * Total Submissions: 182K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 *
 * 说明:
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	//双指针从后往前
	p1 := m - 1
	p2 := n - 1
	i := m + n - 1
	for i >= 0 && p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
		i--
	}
	// nums2 有可能会剩下部分
	for i := 0; i <= p2; i++{
		nums1[i] = nums2[i]
	}
}
// @lc code=end

