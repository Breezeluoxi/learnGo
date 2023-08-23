package test

import (
	"fmt"
	"os/user"
	"runtime"
)

func main() {
	var user, error = user.Current()
	if error == nil {
		fmt.Println("hello,", user.Username)
	}
	// 计算机主板的型号
	fmt.Println(runtime.GOARCH)
	// 操作系统的类型
	fmt.Println(runtime.GOOS)
	// 定义数组

}
