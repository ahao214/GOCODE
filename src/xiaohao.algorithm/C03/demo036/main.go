package main

import (
	"fmt"
	"math/rand"
)

//三扇门
func main() {
	//换门遇见天使的此时和不换门遇见天使的次数
	changeAngelCount, unchangeAngelCount := 0, 0
	for i := 0; i < 1000000; i++ {
		//门的总数
		doors := []int{0, 1, 2}
		//天使门和选中的门
		angelDoor, selectedDoor := rand.Intn(3), rand.Intn(3)
		//上帝移除一扇恶魔门
		for j := 0; j < len(doors); j++ {
			if doors[j] != selectedDoor && doors[j] != angelDoor {
				doors = append(doors[:j], doors[j+1:]...)
				break
			}
		}
		//统计
		if selectedDoor == angelDoor {
			unchangeAngelCount++
		} else {
			changeAngelCount++
		}
	}
	fmt.Println("不换门遇见天使次数：", unchangeAngelCount, "比例：", float32(unchangeAngelCount)/1000000)
	fmt.Println("换门遇见天使次数：", changeAngelCount, "比例：", float32(changeAngelCount)/1000000)
}
