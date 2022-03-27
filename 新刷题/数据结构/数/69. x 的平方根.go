package 数

// mySqrt 获取平方根。 (二分法1)
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

// mySqrt 获取平方根。 (二分法2)
func mySqrt(x int) int {
	left, right := 0, 1<<16
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == x {
			left = mid + 1
		} else if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

// mySqrt 获取平方根。 (二分法3 --- 错误解法)
func mySqrt(x int) int {
	left, right := 0, 1<<16
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == x {
			right = mid - 1
		} else if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 到达边界后，返回 right，此时结果不正确。 (比如左右指针相等时，left=3、right=3 此时 x=9，下一步会变为 left=3、right=2)
	// 到达边界后，返回 left，此时结果不正确。  (比如左右指针相等时，left=3、right=3 此时 x=8，下一步会变为 left=3、right=2)
	// 因此，square == x 时，只能操作 left，让 left = mid + 1，此时:
	// 到达边界后，返回 right，此时结果正确。  (比如左右指针相等时，left=3、right=3 此时 x=9，下一步会变为 left=3、right=4)
	// 到达边界后，返回 right，此时结果正确。  (比如左右指针相等时，left=3、right=3 此时 x=8，下一步会变为 left=3、right=2)
	return left
}

// 猜测点:
//		1. for left <= right 时，下面的判断条件不能是 right = mid or left = mid，必须对 mid 有个偏移，比如 +1、-1。
//      2. 对于有向下取整逻辑的二分，最后只能返回 right，在相等时，应该让 left = mid + 1，尝试扩大。
//		3. 如果题目有向上取整逻辑，那最后只能返回 left， 在相等时，应该让 right = mid - 1，尝试缩小。
