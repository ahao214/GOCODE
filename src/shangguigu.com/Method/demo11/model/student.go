package model

type Student struct{
	Name string
	Age int
}

//当结构体首字母是小写，只能在model使用
//通过工厂模式来解决
type teacher struct{
	Name string
	age int
}

func NewTeacher(n string,a int)*teacher{
	return &teacher{
		Name:n,
		age:a,
	}
}

//如果teacher里面的Age首字母小写，则在其他包不可以直接访问
//我们可以提供一个方法
func(t *teacher) GetAge()int{
	return t.age
}
