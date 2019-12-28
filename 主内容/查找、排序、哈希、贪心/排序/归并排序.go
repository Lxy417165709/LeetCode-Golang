package main

// 归并排序入口
func mergeSort(nums []int) {
	mergeSortExec(nums, 0, len(nums)-1)
}

// 对区间[l, r]执行归并排序
func mergeSortExec(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSortExec(nums, l, mid)
	mergeSortExec(nums, mid+1, r)
	merge(nums, l, mid, mid+1, r)
}

// 对区间 [l1, r1] 与 [l2, r2] 执行合并
func merge(nums []int, l1, r1, l2, r2 int) {
	basePosition := l1
	tmpArr := []int{}
	for l1 <= r1 && l2 <= r2 {
		if nums[l1] <= nums[l2] {
			tmpArr = append(tmpArr, nums[l1])
			l1++
		} else {
			tmpArr = append(tmpArr, nums[l2])
			l2++
		}
	}
	for l1 <= r1 {
		tmpArr = append(tmpArr, nums[l1])
		l1++
	}
	for l2 <= r2 {
		tmpArr = append(tmpArr, nums[l2])
		l2++
	}
	for i := 0; i < len(tmpArr); i++ {
		nums[basePosition+i] = tmpArr[i]
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/sort-an-array/
*/
