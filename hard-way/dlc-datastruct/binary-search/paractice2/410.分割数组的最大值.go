/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 *
 * https://leetcode-cn.com/problems/split-array-largest-sum/description/
 *
 * algorithms
 * Hard (39.57%)
 * Likes:    87
 * Dislikes: 0
 * Total Accepted:    3.5K
 * Total Submissions: 8.7K
 * Testcase Example:  '[7,2,5,10,8]\n2'
 *
 * 给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。
 *
 * 注意:
 * 数组长度 n 满足以下条件:
 *
 *
 * 1 ≤ n ≤ 1000
 * 1 ≤ m ≤ min(50, n)
 *
 *
 * 示例:
 *
 *
 * 输入:
 * nums = [7,2,5,10,8]
 * m = 2
 *
 * 输出:
 * 18
 *
 * 解释:
 * 一共有四种方法将nums分割为2个子数组。
 * 其中最好的方式是将其分为[7,2,5] 和 [10,8]，
 * 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
 *
 *
 */

// @lc code=start

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func split(nums []int, largetSum int) int {
	pieces := 1
	sum := 0
	for _, num := range nums {
		if (sum + num) > largetSum {
			pieces++
			sum = num
		} else {
			sum += num
		}
	}
	return pieces
}

func splitArray(nums []int, m int) int {
	// 二分 最少能切成1份 最多能切成len(nums)份
	// low: max(nums) high:sum(high)

	max, sum := 0, 0
	for _, num := range nums {
		max = maxInt(max, num)
		sum += num
	}
	low := max
	high := sum

	for low < high {
		mid := (low + high) / 2
		pieces := split(nums, mid)
		if pieces > m {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// @lc code=end

