/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (60.41%)
 * Likes:    284
 * Dislikes: 0
 * Total Accepted:    21.7K
 * Total Submissions: 35.4K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和
 * n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 *
 * 示例 1:
 *
 * 输入: [1,3,4,2,2]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [3,1,3,4,2]
 * 输出: 3
 *
 *
 * 说明：
 *
 *
 * 不能更改原数组（假设数组是只读的）。
 * 只能使用额外的 O(1) 的空间。
 * 时间复杂度小于 O(n^2) 。
 * 数组中只有一个重复的数字，但它可能不止重复出现一次。
 *
 *
 */

// @lc code=start
func findDuplicate(nums []int) int {
	// 二分 分为两个部分之后，找中值
	// 遍历数组如果比 中值小的数的数量超过了中值，则说明出现在low-mid之前
	low := 0
	hight := len(nums) - 1

	for low < hight {
		mid := (low + hight) / 2
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			low = mid + 1
		} else {
			hight = mid
		}
	}
	return low

}

// @lc code=end

