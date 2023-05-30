package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	// 第1个字段：秒（0 - 59）
	// 第2个字段：分钟（0 - 59）
	// 第3个字段：小时（0 - 23）
	// 第4个字段：日（1 - 31）
	// 第5个字段：月（1 - 12）
	// 第6个字段：星期（0 - 6）（星期日=0）
	id, err := c.AddFunc("0/5 * * * * ?", func() { // 每分钟的0秒运行
		fmt.Println("每5s执行一次: ", time.Now())
	})
	if err != nil {
		fmt.Println("Error scheduling job: ", err)
	}

	_, err = c.AddFunc("0 0 * * * ?", func() { // 每小时的0分0秒运行
		fmt.Println("每小时执行一次: ", time.Now())
	})
	if err != nil {
		fmt.Println("Error scheduling job: ", err)
	}

	c.Start()

	// 等待一段时间，然后删除原来的任务，添加一个新的任务
	time.Sleep(1 * time.Minute)
	c.Remove(id)

	_, err = c.AddFunc("0/10 * * * * ?", func() {
		fmt.Println("现在变成每10s执行一次: ", time.Now())
	})
	if err != nil {
		fmt.Println("Error scheduling job: ", err)
	}

	// 让主线程不退出，否则所有的定时任务都不会运行
	select {}
}
