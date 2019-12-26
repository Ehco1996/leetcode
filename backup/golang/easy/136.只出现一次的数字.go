package easy

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode-cn.com/problems/single-number/description/
 *
 * algorithms
 * Easy (62.58%)
 * Likes:    817
 * Dislikes: 0
 * Total Accepted:    95.6K
 * Total Submissions: 151.5K
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 *
 * 输入: [2,2,1]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入: [4,1,2,1,2]
 * 输出: 4
 *
 */
func singleNumber(nums []int) int {
	// 交换律：a ^ b ^ c <=> a ^ c ^ b
	// 任何数于0异或为任何数 0 ^ n => n
	// 相同的数异或为0: n ^ n => 0
	a := 0
	for _, num := range nums {
		a = a ^ num
	}
	return a
}
