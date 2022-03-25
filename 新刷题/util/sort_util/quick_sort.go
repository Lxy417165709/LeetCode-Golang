package sort_util

import "github.com/Lxy417165709/LeetCode-Golang/新刷题/util/struct_util"

// QuickSort 快速排序。
func QuickSort(nums []int) {
	// 1. 元素小于2的数组，直接返回。
	if len(nums) < 2 {
		return
	}

	// 2. 分区
	index := Partition(nums)

	// 3. 递归。
	QuickSort(nums[:index])
	QuickSort(nums[index+1:])
}

// Partition 分区。
func Partition(nums []int) int {
	// 1. 保证: 左边的数 <= 基准值 <= 右边的数。
	reference := nums[0]
	left, right := 0, len(nums)-1
	for left <= right {
		// 1.1 自左向右，获取第一个大于基准值的索引。
		for left <= right && nums[left] <= reference {
			left++
		}

		// 1.2 自右向左，获取第一个小于基准值的索引。
		for left <= right && nums[right] >= reference {
			right--
		}

		// 1.3 保证两个索引存在。 (如果其中一个不存在，此时必然是 left = right+1)
		if left <= right {
			// 1.3.1 交换。
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// 2. 摆正基准值位置。
	// 为什么这里一定要是 right 呢? 因为 right 最终指向的数一定小于等于基准值。场景: [1 2 9 6 8]
	nums[0], nums[right] = nums[right], nums[0]

	// 3. 返回基准值位置。
	return right
}

// HeapSort 堆排序。
func HeapSort(nums []int) []int {
	// 1. 建堆。
	heap := struct_util.NewMyHeap(len(nums), func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	// 2. 插入数组元素。
	for _, num := range nums {
		heap.Push(num)
	}

	// 3. 获取排序结果。
	result := make([]int, 0)
	for _, obj := range heap.Sort() {
		result = append(result, obj.(int))
	}

	// 4. 返回。
	return result
}
