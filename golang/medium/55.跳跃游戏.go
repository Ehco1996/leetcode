package medium

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (35.86%)
 * Likes:    314
 * Dislikes: 0
 * Total Accepted:    33K
 * Total Submissions: 91.2K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 *
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个位置。
 *
 * 示例 1:
 *
 * 输入: [2,3,1,1,4]
 * 输出: true
 * 解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,1,0,4]
 * 输出: false
 * 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
 *
 *
 */
func canJump(nums []int) bool {
	// 贪心
	// 在当前回合 选择跳的最远的：（当前步数 + 下一格的步数）
	idx := 0
	for idx < len(nums) && nums[idx] != 0 {
		num := nums[idx]
		if num == 0 {
			break
		}
		max := idx + num
		if max > len(nums)-1 {
			return true
		}
		for i := 1; i < num; i++ {
			if i+nums[idx+i] > num+nums[max] {
				max = idx + i
			}
		}
		idx = max
	}
	return idx >= len(nums)-1
}
