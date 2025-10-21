package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("空数组：", a)

	a[4] = 100
	fmt.Println("设置后：", a)
	fmt.Println("获取：", a[4])

	fmt.Println("长度：", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("声明：", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("声明：", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("索引设置：", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("二维数组：", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("二维数组：", twoD)
}
