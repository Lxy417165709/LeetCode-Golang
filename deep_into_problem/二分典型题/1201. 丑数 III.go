package 二分典型题

const inf = 1000000000000

func nthUglyNumber(n int, a int, b int, c int) int {
	left, right := 1, inf
	for left <= right {
		mid := (left + right) / 2
		count := getCountOfUglyNumberLessThanX(a, b, c, mid)
		if count == n {
			right = mid - 1
		} else if count > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := lcm(arr[:len(arr)-1]...), arr[len(arr)-1]
	return a / gcd(a, b) * b
}

func getCountOfUglyNumberLessThanX(a, b, c, x int) int {
	countOfUglyNumberHasFactorA := x / lcm(a)
	countOfUglyNumberHasFactorB := x / lcm(b)
	countOfUglyNumberHasFactorC := x / lcm(c)

	countOfUglyNumberHasFactorAB := x / lcm(a, b)
	countOfUglyNumberHasFactorBC := x / lcm(b, c)
	countOfUglyNumberHasFactorAC := x / lcm(a, c)

	countOfUglyNumberHasFactorABC := x / lcm(a, b, c)

	return countOfUglyNumberHasFactorA + countOfUglyNumberHasFactorB + countOfUglyNumberHasFactorC - countOfUglyNumberHasFactorAB - countOfUglyNumberHasFactorBC - countOfUglyNumberHasFactorAC + countOfUglyNumberHasFactorABC
}
