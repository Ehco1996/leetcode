func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	// 动态规划 前i天的最大收益 = max(前i-1天的最大收益，第i天的价格-前i-1天中的最小价格)
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfile := 0

	for _, price := range prices {
		minPrice = min(minPrice, price)
		maxProfile = max(maxProfile, price-minPrice)
	}
	return maxProfile
}