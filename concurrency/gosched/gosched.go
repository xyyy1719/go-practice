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
	// 通知Go Runtime运行其他goroutine，但并不会block主程序的运行
	// 一旦主程序运行完毕，其他goroutine的程序也不会继续执行下去
	runtime.Gosched()
	fmt.Println("bye")
}
