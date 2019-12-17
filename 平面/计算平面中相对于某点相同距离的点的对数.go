package main

/*
    给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，
    其中 i 和 j 之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
*/

// 计算平面中相同距离点的对数 (排列数，不是组合)
func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		sameDistancePointCount := make(map[int]int)
		for t := 0; t < len(points); t++ {
			dis := distance(points[i], points[t])
			sameDistancePointCount[dis]++

			// 每增加一个等距离点，可选的回旋镖排列 + 2*(sameDistancePointCount[dis]-1) 种
			// 如果不考虑元组顺序，那么就不需要 * 2
			ans += 2 * (sameDistancePointCount[dis] - 1)
		}
	}
	return ans
}

// 距离 (不开方)
func distance(a []int, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

/*
	题目链接:
		https://leetcode-cn.com/problems/number-of-boomerangs/		 回旋镖的数量
*/
/*
	总结
	1. 题目本质就是找平面中相对于某个点相同距离的点有多少对
*/