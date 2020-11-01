package main

import "math/rand"

// 数组
func getSum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func getPrefixMin(arr []int) []int {
	prefixMin := make([]int, len(arr))
	prefixMin[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixMin[i] = min(prefixMin[i-1], arr[i])
	}
	return prefixMin
}

func getSuffixMin(arr []int) []int {
	suffixMin := make([]int, len(arr))
	suffixMin[len(arr)-1] = arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], arr[i])
	}
	return suffixMin
}


func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}


// 升序数组
func isInArray(sortedArray []int, ref int) bool {
	left, right := 0, len(sortedArray)-1
	for left <= right {
		mid := (left + right) / 2
		if sortedArray[mid] == ref {
			return true
		}
		if sortedArray[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func getCountOfNum(arr []int,ref int) int{
	indexOfFirstEqual := getIndexOfFirstEqual(arr,ref)
	indexOfLastEqual := getIndexOfLastEqual(arr,ref)
	if indexOfFirstEqual==-1{
		return 0
	}
	return indexOfLastEqual-indexOfFirstEqual+1
}
func getIndexOfFirstEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	isExist := false
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == ref {
			isExist = true
			right = mid - 1
		} else {
			if arr[mid] > ref {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if isExist {
		return left
	}
	return -1
}
func getIndexOfLastEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	isExist := false
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == ref {
			isExist = true
			left = mid + 1
		} else {
			if arr[mid] > ref {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	if isExist {
		return right
	}
	return -1
}

func getIndexOfFirstGreater(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func getIndexOfFirstGreaterOrEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func getIndexOfLastSmall(arr []int, ref int) int {
	return getIndexOfFirstGreaterOrEqual(arr, ref) - 1
}
func getIndexOfLastSmallOrEqual(arr []int, ref int) int {
	return getIndexOfFirstGreater(arr, ref) - 1
}

func makeKthSmallToRightPlace(nums []int, KthSmall int) {
	// 例子: A 的 rightPlace 就是：A处于某个位置，这个位置左边的元素都小于等于它，右边的元素都大于等于它。

	l, r := 0, len(nums)-1
	for l <= r {
		index := randomPartition(nums, l, r)
		XthSmall := index + 1
		if XthSmall == KthSmall {
			return
		} else {
			if XthSmall > KthSmall {
				r = index - 1
			} else {
				l = index + 1
			}
		}
	}
	// 到达这里表示:  原数组不存在第K小
}
func makeKthBigToRightPlace(nums []int, KthBig int) {
	makeKthSmallToRightPlace(nums, len(nums)-KthBig+1)
}

func randomPartition(nums []int, l int, r int) int {
	// 返回值是一个索引，假如记做 index
	// 则 nums[index] 位于 rightPlace
	// 同 partition函数

	upsetArrayRandomly(nums, l, r)
	return partition(nums, l, r)
}
func upsetArrayRandomly(nums []int, l int, r int) {
	randomIndex := rand.Intn(r-l+1) + l
	nums[randomIndex], nums[l] = nums[l], nums[randomIndex]
}
func partition(nums []int, l int, r int) int {
	refIndex := l // 这个索引取 (l+r)/2 也可以
	for l <= r {
		for l <= r && nums[l] <= nums[refIndex] {
			l++
		}
		for l <= r && nums[r] >= nums[refIndex] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[refIndex], nums[r] = nums[r], nums[refIndex]
	return r
}
