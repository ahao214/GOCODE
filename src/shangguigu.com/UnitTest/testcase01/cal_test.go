package testcase01

import (
	"testing"
)

//编写测试用例，去测试add是否正确
func TestAdd(t *testing.T) {
	//调用
	res := add(10)
	if res != 55 {
		t.Fatalf("执行add()错误，期望值是：%d,实际值是：%v\n", 55, res)
	}
	//如果正确,输出日志
	t.Logf("add()函数，执行正确")
}
