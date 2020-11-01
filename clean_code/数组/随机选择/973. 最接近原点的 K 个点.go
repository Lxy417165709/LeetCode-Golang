package 随机选择

// ------------------------ 随机选择算法 ------------------------
func kClosest(points [][]int, K int) [][]int {
	left, right := 0, len(points)-1
	for {
		order := partition(points, left, right) + 1
		if order == K {
			return points[:K]
		}
		if order > K {
			right = order - 2
		} else {
			left = order
		}
	}
}

func partition(points [][]int, left, right int) int {
	guard := left
	guardDis := getDistanceToOrigin(points[guard])
	for left <= right {
		for left <= right && getDistanceToOrigin(points[left]) <= guardDis {
			left++
		}
		for left <= right && getDistanceToOrigin(points[right]) >= guardDis {
			right--
		}
		if left <= right {
			points[left], points[right] = points[right], points[left]
		}
	}
	points[guard], points[right] = points[right], points[guard]
	return right
}

func getDistanceToOrigin(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

/*
	题目链接: https://leetcode-cn.com/problems/k-closest-points-to-origin/
	总结:
		1. 其实可以提取出一个模板来解决随机选择问题！！
*/
