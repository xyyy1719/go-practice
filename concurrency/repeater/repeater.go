// 在30秒内重复键盘输入，30秒后自动退出
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second) //让主程序等待30秒
	fmt.Println("Time out.")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
