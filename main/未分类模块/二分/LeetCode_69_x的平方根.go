package 二分

import "fmt"

const (
	inf float64 = 1e10  // 无穷 (如果太小就无法计算较大数的平方根，从而出现差错)
	eps float64 = 1e-10 // 精度 (如果太大就会不够精确，从而出现差错)(精度越大，计算时间越长)
						// (eps太难掌握，所以下面没使用)
)

// 二分 求平方根
func sqrt(x float64) (float64, error) {
	if x < 0 {
		err := fmt.Errorf("负数不支持求平方根")
		return 0, err
	}

	l, r := 0.0, inf

	// 由于eps太难掌握，这里直接循环100次，这样精度已经很高了
	for i := 0; i < 100; i++ {
		mid := l + (r - l) / 2
		evaluationValue := mid * mid
		if evaluationValue > x {
			r = mid
		} else {
			l = mid
		}
	}
	return (l + r) / 2, nil
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

/*
	题目链接:
		https://leetcode-cn.com/problems/sqrtx/submissions/
*/
