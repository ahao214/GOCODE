package main

//121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	buy := 0
	benefit := 0
	for sell := 0; sell < len(prices); sell++ {
		if prices[buy] > prices[sell] {
			buy = sell
		}
		benefit = max(prices[sell]-prices[buy], benefit)
	}
	return benefit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
