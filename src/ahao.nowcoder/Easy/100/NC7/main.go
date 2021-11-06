package main

import "math"

/**
 * NC7 买卖股票的最好时机
 * @param prices int整型一维数组
 * @return int整型
 */
func maxProfit(prices []int) int {
	minprice := math.MaxInt64
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}
