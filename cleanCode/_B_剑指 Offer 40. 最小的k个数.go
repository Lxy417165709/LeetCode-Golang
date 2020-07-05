package main

func getLeastNumbers(arr []int, k int) []int {
	selectSmallKth(arr,k)
	return arr[:k]
}

func selectSmallKth(nums []int, k int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		index := randomPartition(nums, l, r)
		if index+1 == k {
			return nums[index]
		} else {
			if index+1 > k {
				r = index - 1
			} else {
				l = index + 1
			}
		}
	}
	return -100000000 // 表示没找到 (k非法了)
}

// 选择第k大的数 (重复的数次序不等)
func selectBigKth(nums []int, k int) int {
	// 第k大 == 第 len(nums)-k+1 小
	return selectSmallKth(nums, len(nums)-k+1)
}

// 随机划分 (l，r在合法范围内)
// 作用: 经过划分后，使得 元素x左边 <= 元素x <= 元素x右边，返回元素x此时在数组中的索引。
func randomPartition(nums []int, l int, r int) int {
	randomIndex := rand.Intn(r-l+1) + l                     // 选择随机索引
	nums[randomIndex], nums[l] = nums[l], nums[randomIndex] // 打乱数组
	guardIndex := l
	for l <= r {
		for l <= r && nums[l] <= nums[guardIndex] {
			l++
		}
		for l <= r && nums[r] >= nums[guardIndex] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[guardIndex], nums[r] = nums[r], nums[guardIndex]
	return r
}



/*
	总结
		1. 这题我是使用随机选择算法AC的，题解还有用排序、堆解决的。
*/
