package main

import (
	"fmt"
)

type Camera struct {
}

func (c *Camera) TakePicture() string {
	return "拍照"
}

type Phone struct {
}

func (p *Phone) Call() string {
	return "响铃"
}

type Video struct {
}

func (v *Video) TakeVideo() string {
	return "拍视频"
}

type CameraPhone struct {
	Camera
	Phone
	Video
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("我们的新款拍照手机有多种功能")
	fmt.Println("打开了相机：", cp.TakePicture())
	fmt.Println("来电话了：", cp.Call())
	fmt.Println("开始拍视频了：", cp.TakeVideo())
}
