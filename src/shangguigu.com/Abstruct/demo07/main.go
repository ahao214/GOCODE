package main

import(
	"fmt"
)

type Goods struct{
	Name string
	Price float64
}

type Brand struct{
	Name string
	Address string
}

type TV struct{
	Goods
	Brand
}

type TV2 struct{
	*Goods
	*Brand
}


type Monster struct{
	Name string
	Age int
}

type E struct{
	Monster
	int 
	n int
}

func main(){
	tv:=TV{
		Goods{"电视001",3000.09},
		Brand{"海尔","青岛"},
	}

	tv2:=TV{
		Goods{
			Name:"电视002",
			Price:3500.99,
		},
		Brand{
			Name:"松下",
			Address:"日本",
		},
	}

	fmt.Println("TV1=",tv)
	fmt.Println("TV2=",tv2)

	tv3:=TV2{
		&Goods{
			Name:"电视003",
			Price:3500.99,
		},
		&Brand{
			Name:"新飞",
			Address:"河南",
		},
	}
	fmt.Println("TV3=",*tv3.Goods,*tv3.Brand)

	var e E
	e.Name="牛魔王"
	e.Age=300
	e.int=20
	e.n=30
	fmt.Println("e=",e)

}