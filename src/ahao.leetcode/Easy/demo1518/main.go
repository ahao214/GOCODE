package main

//1518. 换酒问题
func numWaterBottles(numBottles int, numExchange int) int {
	//喝的酒瓶数量
	result := numBottles
	//剩余的数量
	rest := numBottles
	for rest >= numExchange {
		newBottles := rest / numExchange
		result += newBottles
		mod := rest % numExchange
		rest = mod + newBottles
	}
	return result
}
