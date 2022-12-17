package main

import (
	"github.com/lutasam/check_in_sys/biz/cronjob"
	"github.com/robfig/cron/v3"
)

func InitCronJob() {
	c := cron.New()
	//// demo
	//_, err := c.AddFunc("@every 5s", cronjob.TestCron)
	//if err != nil {
	//	panic(err)
	//}

	// 更新每日用户打卡状态
	_, err := c.AddFunc("@daily", cronjob.CleanUserRecordStatus)
	if err != nil {
		panic(err)
	}

	c.Start()
}
