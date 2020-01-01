package main

func lemonadeChange(bills []int) bool {
	money := make([]int, 2)
	// money[0]: 5块的张数
	// money[1]: 10块的张数

	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			money[0]++
		case 10:
			if money[0] == 0 {
				return false
			}
			money[0]--
			money[1]++
		case 20:
			// 方案(1)
			if money[0] >= 1 && money[1] >= 1 {
				money[0]--
				money[1]--
				continue
			}
			// 方案(2)
			if money[0] >= 3 {
				money[0] -= 3
				continue
			}
			return false
		}
	}
	return true
}
/*
	总结
	1. 这题在遇到20块的时候有2种方法找零:
			(1) 1张10元 + 1张5元
			(2) 3张5元
		显然，我们应该优先选择方案(1)，因为10块可以由2张5块组成。只有在方案(1)不满足时，才选方案(2)
		如果还不满足，那么就没办法了，直接返回false。
*/
