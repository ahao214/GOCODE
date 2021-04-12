package main

import (
	"ahao.bbs/server/common"
)

func StartOn() {
	if !common.IsProd() {
		return
	}
	//开启定时任务
	startSchedule()
}
