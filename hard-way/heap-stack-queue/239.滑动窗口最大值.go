/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (42.09%)
 * Likes:    149
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 37.9K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 *
 * 返回滑动窗口中的最大值。
 *
 *
 *
 * 示例:
 *
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7]
 * 解释:
 *
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 *
 *
 *
 * 提示：
 *
 * 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
 *
 *
 *
 * 进阶：
 *
 * 你能在线性时间复杂度内解决此题吗？
 *
 */

// @lc code=start
// Dequeue 单调队列 队列头是最大的当前最大的元素
type Dequeue struct {
	nums []int
}

func (dq *Dequeue) max() int {
	return dq.nums[0]
}

func (dq *Dequeue) push(x int) {
	// 添加元素之前需要删除所有比当前元素小的值
	for len(dq.nums) > 0 && dq.nums[len(dq.nums)-1] < x {
		dq.nums = dq.nums[:len(dq.nums)-1]
	}
	dq.nums = append(dq.nums, x)
}

func (dq *Dequeue) pop(x int) {
	// 删除滑动窗口外的值（这个值有可能在push的时候就被删了）
	if len(dq.nums) > 0 && dq.nums[0] == x {
		dq.nums = dq.nums[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	window := Dequeue{}

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			// 先把window填满
			window.push(nums[i])
		} else {
			// window 向前移动
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
	}
	return res
}

// @lc code=end

