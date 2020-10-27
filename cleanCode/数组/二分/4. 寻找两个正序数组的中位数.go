package 二分

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	wholeLength := len(nums1) + len(nums2)
	leftMidNum := getKthSmall(nums1, nums2, (wholeLength+1)/2)
	rightMidNum := getKthSmall(nums1, nums2, (wholeLength+2)/2)
	return float64(leftMidNum+rightMidNum) / 2
}

func getKthSmall(nums1 []int, nums2 []int, k int) int {
	if len(nums2) > len(nums1) {
		return getKthSmall(nums2, nums1, k)
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		if nums1[0] >= nums2[0] {
			return nums2[0]
		}
		return nums1[0]
	}
	refIndex := min(len(nums1)-1, len(nums2)-1, k/2-1)
	if nums1[refIndex] >= nums2[refIndex] {
		return getKthSmall(nums1, nums2[refIndex+1:], k-refIndex-1)
	} else {
		return getKthSmall(nums1[refIndex+1:], nums2, k-refIndex-1)
	}
}

func min(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], min(nums[1:]...)
	if a > b {
		return b
	}
	return a
}

/*
	总结：
	1. 这题本质上就是求第 K 小，进阶之处在于，我们需要在两个数组中找第 K 小。
	2. 二分法最重要的就是找基准！！！ 代码 refIndex := min(len(nums1)-1, len(nums2)-1, k/2-1) 就是基准获取。
*/
