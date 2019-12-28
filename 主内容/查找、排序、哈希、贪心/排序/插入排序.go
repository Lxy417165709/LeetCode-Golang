package main

import (
	"fmt"
	"math/rand"
)

// 稳定的插入排序
func insertSort1(nums []int) {
	for i := 0; i < len(nums); i++ {
		for t := i; t >= 1; t-- {
			if nums[t] >= nums[t-1] {
				break
			}
			nums[t], nums[t-1] = nums[t-1], nums[t]
		}
	}
}

// 稳定的插入排序
func insertSort2(nums []int) {
	for i := 0; i < len(nums); i++ {
		insertPosition,insertNumber := i,nums[i]
		for insertPosition >= 1 && nums[insertPosition-1] > insertNumber {
			nums[insertPosition] = nums[insertPosition-1]
			insertPosition--
		}
		nums[insertPosition] = insertNumber
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
	arr := generate(4000)
	insertSort2(arr)
	fmt.Println(isSequential(arr))
	fmt.Println(arr)
	arr = generate(4000)
	insertSort2(arr)
	fmt.Println(isSequential(arr))
}
