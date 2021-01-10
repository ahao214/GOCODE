package testcase01

import "testing"

func TestGetSub(t *testing.T) {
	//调用
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("执行getSub()错误，期望值是：%d,实际值是：%v\n", 7, res)
	}
	//如果正确,输出日志
	t.Logf("getSub()函数，执行正确")
}
