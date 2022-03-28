package 一维数组

import (
	"fmt"
	"math"
)

// findMedianSortedArrays 寻找两个有序数组的中位数。 (可以转换为找第 k 大的问题)
// 两个数组元素总长度 length:
// 1. 偶数: 找 length/2, length/2 + 1
// 2. 奇数: 找 length/2 + 1。
// 汇总为: math.Ceil(float64(length)/2.0), int(float64(length)/2.0)+1
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	return float64(
		getKthSmallInTwoSortedArray(nums1, nums2, int(math.Ceil(float64(length)/2.0)))+
			getKthSmallInTwoSortedArray(nums1, nums2, int(float64(length)/2.0)+1),
	) / 2.0
}

// getKthSmallInTwoSortedArray 获取两个有序数组中的第k小。 (k从1开始)
func getKthSmallInTwoSortedArray(nums1 []int, nums2 []int, k int) int {
	fmt.Println(nums1, nums2, k)

	// 1. 递归出口。
	if len(nums1) == 0 {
		return getKthSmallInSortedArray(nums2, k)
	}
	if len(nums2) == 0 {
		return getKthSmallInSortedArray(nums1, k)
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	// 2. 获取两个数组的第 k/2 小。 (如果数组不存在 k/2 小，此时取数组末尾元素)
	orderInNums1 := min(k/2, len(nums1)) // 元素在数组1的序号。 (序号从1开始)
	orderInNums2 := min(k/2, len(nums2)) // 元素在数组2的序号。 (序号从1开始)
	targetInNums1 := getKthSmallInSortedArray(nums1, orderInNums1)
	targetInNums2 := getKthSmallInSortedArray(nums2, orderInNums2)

	// 3. 二分递归。
	if targetInNums1 == targetInNums2 {
		return getKthSmallInTwoSortedArray(nums1[orderInNums1:], nums2, k-orderInNums1)
	} else if targetInNums1 > targetInNums2 {
		return getKthSmallInTwoSortedArray(nums1, nums2[orderInNums2:], k-orderInNums2)
	} else {
		return getKthSmallInTwoSortedArray(nums1[orderInNums1:], nums2, k-orderInNums1)
	}

	//if targetInNums1 == targetInNums2 {
	//	return getKthSmallInTwoSortedArray(nums1[orderInNums1:], nums2[orderInNums2-1:], k-orderInNums1-orderInNums2+1)
	//} else if targetInNums1 > targetInNums2 {
	//	return getKthSmallInTwoSortedArray(nums1[orderInNums1-1:], nums2[orderInNums2:], k-orderInNums2-orderInNums1+1)
	//} else {
	//	return getKthSmallInTwoSortedArray(nums1[orderInNums1:], nums2[orderInNums2-1:], k-orderInNums1-orderInNums2+1)
	//}
	// 场景: nums1-[4, 5, 6]、nums2-[1, 2, 3, 6] 寻找第 4 小，按照注释的代码，会找到元素 3，但实际应该是 4。
	// 因为: 5 > 2 并不能保证第 4 小在 nums1-[5, 6] 之中，只能保证第 4 小不在 nums2-[1, 2] 中。
}

// getKthSmallInSortedArray 获取有序数组中的第k小。
func getKthSmallInSortedArray(nums []int, k int) int {
	index := k - 1
	return nums[index]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
