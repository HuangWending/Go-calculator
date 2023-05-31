# Go-calculator
Go计算器程序
package main: 定义了一个名为"main"的包，表示这是一个可执行程序的入口点。

import: 导入了两个包，fmt用于格式化输入输出，strconv用于类型转换。

func main() { ... }: 定义了程序的主函数，所有的代码都在这个函数中执行。

var operator string: 声明了一个名为operator的字符串变量，用于存储操作符。

var num1, num2 float64: 声明了两个名为num1和num2的浮点数变量，用于存储输入的数字。

fmt.Print("请输入第一个数字："): 输出提示消息，要求用户输入第一个数字。

fmt.Scanln(&num1): 从标准输入读取用户输入的第一个数字，并将其存储在num1变量中。

fmt.Print("请输入操作符（+、-、*、/）："): 输出提示消息，要求用户输入操作符。

fmt.Scanln(&operator): 从标准输入读取用户输入的操作符，并将其存储在operator变量中。

fmt.Print("请输入第二个数字："): 输出提示消息，要求用户输入第二个数字。

fmt.Scanln(&num2): 从标准输入读取用户输入的第二个数字，并将其存储在num2变量中。

result := 0.0: 声明并初始化一个名为result的浮点数变量，用于存储运算结果，默认值为0.0。

switch operator { ... }: 根据operator的值进行分支选择。

 case "+": result = num1 + num2: 如果operator的值是"+"，则执行加法运算，并将结果存储在result变量中。

 case "-": result = num1 - num2: 如果operator的值是"-"，则执行减法运算，并将结果存储在result变量中。
 case "*": result = num1 * num2: 如果operator的值是"*"，则执行乘法运算，并将结果存储在result变量中。

 case "/": ...: 如果operator的值是"/"，则执行除法运算，并将结果存储在result变量中。如果num2不等于0，则执行除法运算；否则输出错误信息并终止程序。

fmt.Println("结果：", strconv.FormatFloat(result, 'f', 2, 64)): 输出结果到标准输出。使用strconv.FormatFloat函数将result转换为字符串，并限定小数点后2位。
