/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (37.69%)
 * Likes:    98
 * Dislikes: 0
 * Total Accepted:    5.3K
 * Total Submissions: 14K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于
 * nums[i] 的元素的数量。
 *
 * 示例:
 *
 * 输入: [5,2,6,1]
 * 输出: [2,1,1,0]
 * 解释:
 * 5 的右侧有 2 个更小的元素 (2 和 1).
 * 2 的右侧仅有 1 个更小的元素 (1).
 * 6 的右侧有 1 个更小的元素 (1).
 * 1 的右侧有 0 个更小的元素.
 *
 *
 */

// @lc code=start

func bsInsert(nums []int, target int) ([]int, int) {

	if len(nums) == 0 {
		nums = append(nums, target)
		return nums, 0
	}

	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		fmt.Println("in", mid)
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	nums = append(nums, 0 /* use the zero value of the element type */)
	copy(nums[l+1:], nums[l:])
	nums[l] = target
	return nums, l
}

func countSmaller(nums []int) []int {
	// 从后往前依次取出数字去排序
	// 使用二分查找找到新的数字排在第几位，就是比几个数字大
	// 然后将这个数字插入到该位置

	size := len(nums)
	res := make([]int, size)
	sorted := []int{}

	for i := size - 1; i >= 0; i-- {
		n, idx := bsInsert(sorted, nums[i])
		sorted = n
		res[i] = idx
	}
	return res
}

// @lc code=end

