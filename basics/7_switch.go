package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("输出 ", i, " 为 ")
	switch i {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天是工作日")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("现在是中午之前")
	default:
		fmt.Println("现在是中午之后")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("我是一个布尔值")
		case int:
			fmt.Println("我是一个整数")
		default:
			fmt.Printf("不知道类型 %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("AuroraSEC")
}
