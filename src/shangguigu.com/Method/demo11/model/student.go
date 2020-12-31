package model

type Student struct{
	Name string
	Age int
}

//当结构体首字母是小写，只能在model使用
//通过工厂模式来解决
type teacher struct{
	Name string
	Age int
}

func NewTeacher(n string,a int)*teacher{
	return &teacher{
		Name:n,
		Age:a,
	}
}
