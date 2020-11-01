package 双指针

import "fmt"

func twoSum(nums []int, target int) []int {
	var a [2]int
	m1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m1[nums[i]] = i + 1	// 这+1的原因是map默认值为0,但是0可能是下标
	}
	for i := 0; i < len(nums); i++ {
		if m1[target-nums[i]] != 0 && m1[target-nums[i]] - 1!=i {
			a[0] = i
			a[1] = m1[target-nums[i]] - 1	//这里-1是为了还原下标
		}
	}
	return a[:]
}

func main() {
	fmt.Println(twoSum([]int{4,4},8))
}

/*
	总结
	1. 此方法采用了哈希数据结构
	2. 这题还可以用排序 + two point

*/