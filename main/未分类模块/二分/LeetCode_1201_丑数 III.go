package 二分

import "sort"

// 二分..
func nthUglyNumber(n int, a int, b int, c int) int {
	tmp := []int{a, b, c}
	sort.Ints(tmp)

	l := tmp[0]              // 下界
	r := 2000000000000000000 // 上界
	for l <= r {
		mid := (l + r) / 2
		cnt := count(a, b, c, mid)
		if cnt == n {
			r = mid - 1
		} else {
			if cnt > n {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return l
}

// 返回它是第多少个丑数
// 这里用到了容斥定理
func count(a int, b int, c int, num int) int {
	// 这里还可以优化，cnt_x 都是可以固定得到的值。。
	cnt_a, cnt_b, cnt_c := num/LCM(a), num/LCM(b), num/LCM(c)
	cnt_ab, cnt_ac, cnt_bc := num/LCM(a, b), num/LCM(a, c), num/LCM(b, c)
	cnt_abc := num / LCM(a, b, c)
	return cnt_a + cnt_b + cnt_c - cnt_ab - cnt_ac - cnt_bc + cnt_abc
}


// 求数组所有元素的最小公倍数
func LCM(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	x := LCM(nums[1:]...)
	gcd := GCD(x, nums[0])
	return x / gcd * nums[0]
}

// 求数组所有元素的最大公约数
func GCD(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	x, y := nums[0], GCD(nums[1:]...)
	if y == 0 {
		return x
	}
	if x == 0 {
		return y
	}
	return GCD(y, x%y)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}