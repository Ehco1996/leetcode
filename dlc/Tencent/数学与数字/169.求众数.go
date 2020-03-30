/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (60.19%)
 * Likes:    309
 * Dislikes: 0
 * Total Accepted:    65.9K
 * Total Submissions: 109.1K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在众数。
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */
func majorityElement(nums []int) int {
	// 因为总是存在众数，那么这个数的数量一定是>n/2的
	c := 0
	res := nums[0]
	for _, num := range nums {
		if num == res {
			c += 1
		} else {
			c -= 1
		}
		if c == 0 {
			res = num
			c = 1
		}
	}
	return res
}
