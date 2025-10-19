package main

import "fmt"

func main() {

	// 基础 for 循环写法之一：类似 while 循环
	// 当条件 i <= 3 为真时执行循环体
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的三段式 for 循环
	// 初始化语句；条件判断；递增语句
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// 使用 range 遍历（Go 1.22+ 支持整数 range）
	// 这里 range 3 表示从 0 到 2
	for i := range 3 {
		fmt.Println("范围", i)
	}

	// ④ 无限循环（手动使用 break 跳出）
	for {
		fmt.Println("循环")
		break
	}

	// 结合 continue 控制语句的循环
	// 打印 0~5 之间的奇数
	for n := range 6 {
		if n%2 == 0 { // 若是偶数则跳过
			continue
		}
		fmt.Println(n)
	}
}
