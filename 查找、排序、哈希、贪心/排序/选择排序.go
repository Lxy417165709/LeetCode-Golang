package main

import (
	"fmt"
	"math/rand"
)

// 不稳定的选择排序
func selectSort(nums []int) {
	for i:=0;i<len(nums);i++{
		minIndex,minNumber := i,nums[i]
		for t:=i+1;t<len(nums);t++{
			if minNumber>nums[t]{
				minIndex,minNumber = t,nums[t]
			}
		}
		nums[i],nums[minIndex] = nums[minIndex],nums[i]
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
	selectSort(arr)
	fmt.Println(isSequential(arr))
	arr = generate(4000)
	selectSort(arr)
	fmt.Println(isSequential(arr))
}

