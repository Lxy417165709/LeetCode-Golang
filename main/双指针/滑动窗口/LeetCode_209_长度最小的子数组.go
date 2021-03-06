package 滑动窗口

// 滑动窗口
func minSubArrayLen(s int, nums []int) int {

	minLength := len(nums) + 1

	/* 1. 初始化窗口数据结构，用于记录窗口内的信息 */
	sum := 0
	first, last := 0, 0

	for last < len(nums) {
		/* 2. 把 last 指向的元素加入窗口 */
		sum += nums[last]

		/* 3. 判断当前窗口内的元素是否符合条件 */
		for first < len(nums) && sum >= s {
			// 如果题目求的是最小值，则符合题意条件时进入函数体

			/* 4.a 在这写更新窗口最小值的代码 */
			minLength = min(minLength, last-first+1)

			/* 5. 把 first 指向的元素移出窗口 */
			sum -= nums[first]
			first++
		}
		last++
	}
	if minLength == len(nums)+1 {
		return 0
	}
	return minLength
}

func min(a, b int) int{
	if a > b {
		return b
	}
	return a
}


// 滑动窗口 + 二分
func minSubArrayLen(s int, nums []int) int {
	// sum[i] 表示 nums[:i]的总和
	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum [i] = sum[i-1] + nums[i-1]
	}
	minLength := len(nums) + 1
	first, last := 0, 0
	for last < len(sum) {
		idx := lastLessOrEqual(sum, sum[last]-s)
		// 这里一定要判断idx是否存在
		if idx != -1 {
			first = idx + 1
			minLength = min(minLength, last-first+1)
		}
		last ++
	}
	if minLength == len(nums)+1 {
		return 0
	}
	return minLength
}

// 返回nums中最后一个小于或等于target的索引，没有则为-1		(注意这里的nums和上面函数的nums不一样)
func lastLessOrEqual(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			l = mid + 1
		} else {
			if nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return r
}

/*
	题目链接:
		https://leetcode-cn.com/problems/minimum-size-subarray-sum/submissions/			长度最小的子数组
*/

/*
	总结
	1. 这题和之前的滑动窗口做法是差不多的，只是这次求的是最小值，之前求的是最大值。
	2. 这题还可以使用二分 + 滑动窗口。 (不过二分的复杂度是O(nlogn)，单纯用滑动窗口的话复杂度是O(n))
*/
