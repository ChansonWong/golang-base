package main

import (
	"fmt"
	"golag-base/src/function"
	"golag-base/src/pointer"
	"golag-base/src/variables"
)

func main() {
	fmt.Println("========变量使用说明书=======")
	variables.DeclareVariables()

	fmt.Println("========方法使用说明书=======")
	var totalPrice = function.CalculateBill(3, 10)
	fmt.Println("总价：", totalPrice)
	area, perimeter := function.RectProps(10.8, 5.6)
	fmt.Println("area is", area, "perimeter is", perimeter)

	fmt.Println("========指针如果使用=======")
	pointer.PointerBaseUse()
	pointer.SwapValue()
	pointer.NewFunc()
}
