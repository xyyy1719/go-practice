package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	// 匿名方法
	go func() {
		fmt.Println("Inside a goroutine.")
	}()
	fmt.Println("Outside again.")
	runtime.Gosched() // 通知Go Runtime运行其他goroutine
	fmt.Println("bye")
}
