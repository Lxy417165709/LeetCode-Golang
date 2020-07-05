package main

type RecentCounter struct {
	pingArr []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	r.pingArr = append(r.pingArr, t)
	lastPingTimeIndexThatSatisfiesCondition := len(r.pingArr)-1
	firstPingTimeIndexThatSatisfiesCondition := getIndexOfFirstGreaterOrEqual(r.pingArr, t-3000)
	return lastPingTimeIndexThatSatisfiesCondition-firstPingTimeIndexThatSatisfiesCondition+1
}

func getIndexOfFirstGreaterOrEqual(sortedArr []int, findingNum int) int {
	left, right := 0, len(sortedArr)-1
	for left <= right {
		mid := left + (right-left)/2
		if findingNum == sortedArr[mid] {
			right = mid - 1
		} else {
			if findingNum > sortedArr[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

 /*
 	总结
 		1. 写二分查找，首先需要确定是什么类型的，比如是 第一个大于等于、还是第一个大于...
 */
