package pointer

import "fmt"

func PointerBaseUse() {
	var word = "Hello world"

	// 对字符串取地址, ptr类型为*string
	ptr := &word
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
}

/**
 * 注意*作为左值或者右值的区别
 */
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*a = t
}

func SwapValue() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

/**
 * new关键字如果使用
 */
func NewFunc() {
	str := new(string)
	*str = "Hello World"
	fmt.Println(*str)
}
