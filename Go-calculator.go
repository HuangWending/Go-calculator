package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operator string
	var num1, num2 float64

	fmt.Print("请输入第一个数字：")
	fmt.Scanln(&num1)

	fmt.Print("请输入操作符（+、-、*、/）：")
	fmt.Scanln(&operator)

	fmt.Print("请输入第二个数字：")
	fmt.Scanln(&num2)

	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("错误：除数不能为0")
			return
		}
	default:
		fmt.Println("错误：无效的操作符")
		return
	}

	fmt.Println("结果：", strconv.FormatFloat(result, 'f', 2, 64))
}