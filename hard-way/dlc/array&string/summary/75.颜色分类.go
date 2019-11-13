/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 *
 * https://leetcode-cn.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (53.50%)
 * Likes:    265
 * Dislikes: 0
 * Total Accepted:    41.7K
 * Total Submissions: 77.7K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
 *
 * 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
 *
 * 注意:
 * 不能使用代码库中的排序函数来解决这道题。
 *
 * 示例:
 *
 * 输入: [2,0,2,1,1,0]
 * 输出: [0,0,1,1,2,2]
 *
 * 进阶：
 *
 *
 * 一个直观的解决方案是使用计数排序的两趟扫描算法。
 * 首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
 * 你能想出一个仅使用常数空间的一趟扫描算法吗？
 *
 *
 */

// @lc code=start
func sortColors(nums []int) {
	// 因为只有三种颜色，所以可以用三指针来做
	// 比较通用的解决办法是两次遍历  一次得到所有元素的个数 一次写数组

	p0 := 0
	p1 := len(nums) - 1
	for cur := 0; cur <= p1;{
		if nums[cur]==0{
			nums[cur] ,nums[p0] = nums[p0], nums[cur]
			p0++
			cur++
		}else if nums[cur]==2{
			// 只要把cur交换过去就好,当前指针不用移动了
			nums[cur] ,nums[p1] = nums[p1], nums[cur]
			p1--
        }else{
            cur++
        }
	}
}

// @lc code=end

