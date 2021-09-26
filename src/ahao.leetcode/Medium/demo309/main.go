package main

//309. 最佳买卖股票时机含冷冻期
import (
	"math"
)

// 解法一 DP
func maxProfit309(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]
	buy[1] = max(buy[0], -prices[1])
	sell[1] = max(sell[0], buy[0]+prices[1])
	for i := 2; i < len(prices); i++ {
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], sell[i-2]-prices[i])
	}
	return sell[len(sell)-1]
}

func max(buy int, sell int) int {
	if buy > sell {
		return buy
	}
	return sell
}
