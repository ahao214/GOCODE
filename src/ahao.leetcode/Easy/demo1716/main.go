package main

/*
1716. 计算力扣银行的钱
Hercy 想要为购买第一辆车存钱。他 每天 都往力扣银行里存钱。

最开始，他在周一的时候存入 1 块钱。从周二到周日，他每天都比前一天多存入 1 块钱。在接下来每一个周一，他都会比 前一个周一 多存入 1 块钱。

给你 n ，请你返回在第 n 天结束的时候他在力扣银行总共存了多少块钱。

*/

func totalMoney(n int) (ans int) {
	week, day := 1, 1
	for i := 0; i < n; i++ {
		ans += week + day - 1
		day++
		if day == 8 {
			day = 1
			week++
		}
	}
	return
}

func totalMoney1(n int) (ans int) {
	// 所有完整的周存的钱
	weekNum := n / 7
	firstWeekMoney := (1 + 7) * 7 / 2
	lastWeekMoney := firstWeekMoney + 7*(weekNum-1)
	weekMoney := (firstWeekMoney + lastWeekMoney) * weekNum / 2
	// 剩下的不能构成一个完整的周的天数里存的钱
	dayNum := n % 7
	firstDayMoney := 1 + weekNum
	lastDayMoney := firstDayMoney + dayNum - 1
	dayMoney := (firstDayMoney + lastDayMoney) * dayNum / 2
	return weekMoney + dayMoney
}
