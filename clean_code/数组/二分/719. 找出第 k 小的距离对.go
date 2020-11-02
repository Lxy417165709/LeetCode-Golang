package 二分

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	minDis, maxDis := 0, nums[len(nums)-1]-nums[0]
	for minDis <= maxDis {
		midDis := (maxDis + minDis) / 2
		order := getCountOfLessOrEqualDis(nums, midDis)

		// 当 order == k 时，此时 midDis 可能不在数组里面，所以此时继续向左查找
		// (向右查找是不可以的，因为当 mid == 数组的元素时，此时 order 必然 > k)
		if order == k {
			maxDis = midDis - 1
			continue
		}

		// 当 order > k 时, 说明要向左寻找，反之则向右寻找
		if order > k {
			maxDis = midDis - 1
		} else {
			minDis = midDis + 1
		}
	}

	// 出来时，必然是 minDis > maxDis, 且 midDis 比 maxDis 大 1，且 midDis 一定是数组中的元素。
	return minDis
}

func getCountOfLessOrEqualDis(nums []int, refDis int) int {
	start, count := 0, 0
	for stop := 1; stop < len(nums); stop++ {
		for nums[stop]-nums[start] > refDis {
			start++
		}
		count += stop - start
	}
	return count
}

/*
	总结
	1. 二分依据: 索引二分、值二分。
	2. 二分法注意点: 找到参考点。
	3. A是第 n 小， 此时可以分情况讨论：
		1. A 在数组中，此时有 n 个数小于等于 A
		2. A 不在数组中，此时有 n 个数小于 A
*/
