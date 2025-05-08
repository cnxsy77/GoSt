package main

import (
	"fmt"
	"os"
	"time"
)

/*
	1.2 命令行参数
	os.Args
*/

// 练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。
// 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
// 练习 1.3： 做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6 节讲解了部分 time 包，11.4 节展示了如何写标准测试程序，以得到系统性的性能评测。）
func main() {
	now := time.Now()
	for idx, arg := range os.Args {
		if idx == 0 {
			fmt.Println("os.Args[0]=", arg)
		}
		fmt.Println(idx, "=", arg)
	}
	end := time.Now()
	fmt.Println("运行时间", end.Sub(now))
	// fmt.Println(os.Args[0])
}
