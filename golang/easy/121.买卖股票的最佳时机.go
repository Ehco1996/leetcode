// package main

// import "fmt"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (50.41%)
 * Likes:    490
 * Dislikes: 0
 * Total Accepted:    66.1K
 * Total Submissions: 131K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
 *
 * 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
 *
 * 注意你不能在买入股票前卖出股票。
 *
 * 示例 1:
 *
 * 输入: [7,1,5,3,6,4]
 * 输出: 5
 * 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
 * ⁠    注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
 *
 *
 * 示例 2:
 *
 * 输入: [7,6,4,3,1]
 * 输出: 0
 * 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
 *
 *
 */
func maxProfit(prices []int) int {
	// 动态规划 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
	if len(prices)==0{
		return 0
	}
	minPrice := prices[0]
	maxProfile := 0
	for i := 1; i < len(prices); i++ {
		cur := prices[i]
		if minPrice > cur {
			minPrice = cur
		}
		if cur-minPrice > maxProfile {
			maxProfile = cur - minPrice
		}
	}
	return maxProfile
}

// func main() {
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// }
