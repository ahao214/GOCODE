package main

//123. 买卖股票的最佳时机 III
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	//完成第一笔交易获取到的最大值 减去 再次买入股票价格，得到的最大值
	doneOnceMinusBuyMax := -prices[0]
	//完成第一笔交易获取到的最大值
	doneOnceMax := 0
	for i := 1; i < len(prices); i++ {
		min = minPrice(min, prices[i])
		res = maxPrice(res, doneOnceMinusBuyMax+prices[i])
		doneOnceMax = maxPrice(doneOnceMax, prices[i]-min)
		doneOnceMinusBuyMax = maxPrice(doneOnceMinusBuyMax, doneOnceMax-prices[i])
	}
	return res
}

func maxPrice(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minPrice(a, b int) int {
	if a > b {
		return b
	}
	return a
}
