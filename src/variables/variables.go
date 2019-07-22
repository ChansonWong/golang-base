package variables

import "fmt"

/**
 * 声明变量
 */
func DeclareVariables() {
	// 定义变量：var name type
	var age int
	fmt.Println("my age is", age)

	age = 29 // 赋值
	fmt.Println("my age is", age)

	// 声明变量并赋值 var name type = initialvalue
	var year int = 2019
	fmt.Println("This year is", year)

	// 类型推断，声明变量并赋值时不需要写明类型 var name = initialvalue
	var month = 7
	fmt.Println("This month is", month)

	// 声明多个变量
	var width, height int = 100, 50 // 声明多个变量
	fmt.Println("width is", width, "height is", height)

	// 声明与赋值分开
	var hours, minute int
	fmt.Println("hours is", hours, "minute is", minute)
	hours = 1
	minute = 20
	fmt.Println("new hours is", hours, "new minute is ", minute)

	// 简短声明 name := initialvalue
	name, long := "Risc", 182
	fmt.Println("my name is", name, "long is", long)
}
