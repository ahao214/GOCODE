package testcase02

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//给monster绑定一个Store方法，将monster对象，序列化后保存到文件中
func (this *Monster) Store() bool {
	//序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("序列化失败了：%v\n", err)
		return false
	}

	//保存到文件中
	filePath := "f:/monster.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Printf("write file is err:%v\n", err)
		return false
	}
	return true
}

//给monster绑定一个ReStore方法，将monster对象进行反序列化
func (this *Monster) ReStore() bool {
	//先从文件中，读取序列化的字符串
	filePath := "f:/monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file is err:%v\n", err)
		return false
	}

	//对读取到的data，进行反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("反序列化有错:%v\n", err)
		return false
	}
	return true

}
