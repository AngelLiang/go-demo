package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers() // 启动goroutine来执行printNumbers函数

	// 在主goroutine中执行其他任务
	for i := 0; i < 5; i++ {
		fmt.Println("Main Goroutine:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Child Goroutine:", i)
		time.Sleep(time.Millisecond * 500)
	}
}
