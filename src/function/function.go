package function

/**
 * 普通使用方法
 */
func CalculateBill(price int, no int) int {
	// 商品总价 = 商品单价 * 数量
	var totalPrice = price * no
	// 返回总价
	return totalPrice
}

/**
 * 多值返回
 */
func RectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func RectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	// 不需要明确指定返回值，默认返回 area, perimeter 的值
	return
}
