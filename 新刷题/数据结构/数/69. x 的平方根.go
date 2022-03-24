package 数

func mySqrt(x int) int {
	left, right := 0, 1<<16
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == x {
			return mid
		} else if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
