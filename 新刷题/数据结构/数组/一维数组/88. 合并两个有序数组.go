package 一维数组

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 1. 初始化。
	readIndexOfNums1 := m - 1
	readIndexOfNums2 := n - 1
	writeIndex := m + n - 1

	// 2. 合并。
	for readIndexOfNums1 >= 0 && readIndexOfNums2 >= 0 {
		if nums1[readIndexOfNums1] >= nums2[readIndexOfNums2] {
			nums1[writeIndex] = nums1[readIndexOfNums1]
			readIndexOfNums1--
		} else {
			nums1[writeIndex] = nums2[readIndexOfNums2]
			readIndexOfNums2--
		}

		writeIndex--
	}

	// 3. 写入剩余未比对数。
	for readIndexOfNums1 >= 0 {
		nums1[writeIndex] = nums1[readIndexOfNums1]
		readIndexOfNums1--
		writeIndex--
	}
	for readIndexOfNums2 >= 0 {
		nums1[writeIndex] = nums2[readIndexOfNums2]
		readIndexOfNums2--
		writeIndex--
	}
}
