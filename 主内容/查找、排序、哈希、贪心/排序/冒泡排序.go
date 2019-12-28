package main

import (
	"fmt"
	"math/rand"
)

// 未优化的冒泡排序
func bubbleSort1(nums []int) {
	for i := 0; i < len(nums); i++ {
		for t := 1; t < len(nums)-i; t++ {
			if nums[t-1] > nums[t] {
				nums[t-1], nums[t] = nums[t], nums[t-1]
			}
		}
	}
}

// 优化的冒泡排序
func bubbleSort2(nums []int) {
	for i := 0; i < len(nums); i++ {
		swapFlag := false
		for t := 1; t < len(nums)-i; t++ {
			if nums[t-1] > nums[t] {
				swapFlag = true
				nums[t-1], nums[t] = nums[t], nums[t-1]
			}
		}
		// 没有交换表示数组已经有序了，所以直接退出
		if swapFlag == false {
			break
		}
	}
}

// -------------------- 以下是测试代码 -------------------------
const limit = 500000

// 生成一个随机的数组
func generate(n int) []int {
	sli := make([]int, 0)
	for i := 0; i < n; i++ {
		sli = append(sli, rand.Intn(limit)-limit/2)
	}
	return sli
}

// 判断是否有序(不递减)
func isSequential(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

// 测试
func main() {
	arr := generate(1000)
	bubbleSort1(arr)
	fmt.Println(isSequential(arr))
	arr = generate(1000)
	bubbleSort2(arr)
	fmt.Println(isSequential(arr))
}

/*
	题目链接:
		https://leetcode-cn.com/problems/sort-an-array/		排序数组(会超时)
*/
