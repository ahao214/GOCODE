package testcase02

import (
	"testing"
)

func TestStore(t *testing.T) {
	//先创建一个monster实例
	monster := &Monster{
		Name:  "红孩儿",
		Age:   12,
		Skill: "吐火",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，期望为=%v 实际值为=%v\n", true, res)
	}
	t.Logf("monster.Store()测试正确")
}

func TestReStore(t *testing.T) {
	//先创建一个monster实例,不需要指定字段的值
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，期望为=%v 实际值为=%v\n", true, res)
	}

	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误，期望为=%v 实际值为=%v\n", "红孩儿", monster.Name)
	}

	t.Logf("monster.ReStore()测试正确")

}
