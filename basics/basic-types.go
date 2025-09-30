package main

import "fmt"

func main() {
	// 声明但不初始化 ——> 零值
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("=== Zero Values in Go ===")
	fmt.Printf("int: %d\n", i)
	fmt.Printf("float64: %f\n", f)
	fmt.Printf("bool: %t\n", b)
	fmt.Printf("string: %q\n", s) // 用 %q 可以看到空字符串

	// 显式初始化
	i = 42
	f = 3.1415
	b = true
	s = "GoLang"

	fmt.Println("\n=== After Initialization ===")
	fmt.Printf("int: %d\n", i)
	fmt.Printf("float64: %.4f\n", f)
	fmt.Printf("bool: %t\n", b)
	fmt.Printf("string: %s\n", s)
}
