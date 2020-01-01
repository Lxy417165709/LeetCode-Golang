package 二分

func minEatingSpeed(piles []int, H int) int {
	l, r := 1, 10000000000
	for l <= r {
		mid := l + (r-l)/2
		hour := countHour(piles, mid)	// 测试: 用 mid 速度吃香蕉要花多久
		// 相等则向下试探，即吃慢点。
		if hour == H {
			r = mid - 1
		} else {
			// 超时了，则加快吃的速度
			if hour > H {
				l = mid + 1
			} else {
				// 没超时则继续向下试探，即吃慢点。
				r = mid - 1
			}
		}
	}
	return l	// 此时 l 就是满足条件的最小值
}

// 计算吃完这堆香蕉要花多少小时
func countHour(piles []int, k int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += floor(piles[i], k)
	}
	return hours
}

// 取 a/b 上界
func floor(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}
