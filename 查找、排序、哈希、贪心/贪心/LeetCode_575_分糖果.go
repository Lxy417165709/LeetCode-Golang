package main

func distributeCandies(candies []int) int {
	candyKinds := 0
	candyCount := make(map[int]int)
	for i := 0; i < len(candies); i++ {
		candyCount[candies[i]]++
		if candyCount[candies[i]] == 1 {
			candyKinds++
		}
	}
	if candyKinds > len(candies)>>1 {
		candyKinds = len(candies) >> 1
	}
	return candyKinds
}

/*
	贪心策略:
		让妹妹每一种糖果只拿一个，这样可以保证妹妹糖果的种类最多。
		又为了妹妹和弟弟的糖果均分，假如妹妹手中的糖果大于 总糖果数/2，那就要把妹妹超出的糖果分给弟弟。
*/

/*
	题目链接:
		https://leetcode-cn.com/problems/distribute-candies/solution/tong-guo-bi-jiao-tang-guo-chong-lei-shu-liang-he-t/		分糖果
*/

/*
	总结
	1. 这题目还可以使用排序，排序后获得糖果的种类，这样可以使空间复杂度降为O(1)，但是时间复杂度会升为O(nlogn)
*/
