/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (66.58%)
 * Likes:    134
 * Dislikes: 0
 * Total Accepted:    39.6K
 * Total Submissions: 59.2K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * 说明:
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	s1 := make(map[int]int)
	for _, num := range nums1 {
		s1[num] = num
	}
	res := []int{}
	for _, num := range nums2 {
		if val, ok := s1[num]; ok && val == num {
			res = append(res, num)
			s1[num] = -1
		}
	}
	return res
}

// @lc code=end

