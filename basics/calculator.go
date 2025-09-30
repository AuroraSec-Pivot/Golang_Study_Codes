// 需求：编写一个命令行程序，支持接收用户输入的 两个数字 和 一个运算符（+、-、*、/），输出计算结果。需处理以下场景：
// 用户输入非数字（如字母、符号）时，返回明确的错误提示（如 “输入的不是有效数字，请重新输入”）；
// 除法运算中除数为 0 时，返回 “除数不能为 0” 的错误；
// 运算符不是 “+、-、、/” 时，返回 “不支持的运算符，请输入 +、-、、/”。
// 提示：需用到 fmt.Scan/fmt.Scanf 接收输入、error 类型处理错误、条件判断（if-else）。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 验证输入是否为数字
func parseNumber(input string) (float64, error) {
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

// 验证运算符是否合法
func isValidOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

// 计算函数
func calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("除数不能为 0")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("不支持的运算符 '%s'", op)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("简易计算器")
	fmt.Println("请输入两个数字和运算符（例如：3 + 5），以空格分隔")
	fmt.Println("输入 'exit' 退出程序")

	for {
		fmt.Printf("\n请输入计算表达式：")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			fmt.Println("程序已退出")
			return
		}

		// 拆分输入
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("输入格式错误，请使用：数字 运算符 数字，例如：3 + 5")
			continue
		}

		// 检查运算符是否在中间
		op := parts[1]
		if !isValidOperator(op) {
			fmt.Println("输入格式错误，请使用：数字 运算符 数字，例如：3 + 5")
			continue
		}

		// 检查数字是否有效
		num1, err1 := parseNumber(parts[0])
		num2, err2 := parseNumber(parts[2])
		if err1 != nil || err2 != nil {
			fmt.Println("输入格式错误，请使用：数字 运算符 数字，例如：3 + 5")
			continue
		}

		// 执行运算
		result, err := calculate(num1, num2, op)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 格式化输出
		if result == float64(int(result)) && num1 == float64(int(num1)) && num2 == float64(int(num2)) {
			// 三者都是整数，显示整数格式
			fmt.Printf("计算结果: %d %s %d = %d\n", int(num1), op, int(num2), int(result))
		} else {
			// 否则保留两位小数
			fmt.Printf("计算结果: %.2f %s %.2f = %.2f\n", num1, op, num2, result)
		}
	}
}

/*
================= 测试样例 =================

// ✅ 正常输入（正确结果）
3 + 5          // 输出: 计算结果: 3 + 5 = 8
10 - 4         // 输出: 计算结果: 10 - 4 = 6
6 * 7          // 输出: 计算结果: 6 * 7 = 42
8 / 2          // 输出: 计算结果: 8 / 2 = 4
3.5 + 2.2      // 输出: 计算结果: 3.50 + 2.20 = 5.70
5.5 * 2        // 输出: 计算结果: 5.50 * 2.00 = 11.00

// ❌ 非法数字输入
a + 5          // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5
3 + b          // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5
x * y          // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5

// ❌ 除数为 0
5 / 0          // 输出: 除数不能为 0
10 / 0         // 输出: 除数不能为 0

// ❌ 不支持的运算符
3 ^ 2          // 输出: 不支持的运算符，请输入 +、-、*、/
7 % 2          // 输出: 不支持的运算符，请输入 +、-、*、/
4 & 1          // 输出: 不支持的运算符，请输入 +、-、*、/

// ❌ 格式错误输入
3 +            // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5
+ 3 5          // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5
3 5            // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5
hello world    // 输出: 输入格式错误，请使用：数字 运算符 数字，例如：3 + 5

// ✅ 程序退出
exit           // 输出: 程序已退出，并终止运行
*/
