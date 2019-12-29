package 双指针

import (
	"fmt"
	"sort"
)

// 要求A数组是升序排序的,这个函数的返回值是能组成的 值的二元组
func twoSum(nums []int, sum int) [][]int {
	//sort.Ints(A)	// 注意在函数体外就应该完成排序
	l, r := 0, len(nums)-1
	combinations := make([][]int, 0)
	for l < r {
		tmpSum := nums[l] + nums[r]
		leftAdjustFlag, rightAdjustFlag := false, false

		if sum == tmpSum {
			combinations = append(combinations, []int{nums[l], nums[r]})
			leftAdjustFlag, rightAdjustFlag = true, true
		} else {
			if tmpSum > sum {
				rightAdjustFlag = true
			} else {
				leftAdjustFlag = true
			}
		}
		// 下面的操作是为了保证不出现重复的二元组
		if leftAdjustFlag {
			reference := nums[l]
			for ; l < r && nums[l] == reference; {
				l++
			}
		}
		if rightAdjustFlag {
			reference := nums[r]
			for ; l < r && nums[r] == reference; {
				r--
			}
		}
	}
	return combinations
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	m := make(map[int]int)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			continue
		}
		m[nums[i]] = 1

		slice := twoSum(nums[i+1:], -nums[i])
		for t := 0; t < len(slice); t++ {
			slice[t] = append(slice[t], nums[i])
			temp := make([]int, len(slice[t]))
			copy(temp, slice[t])
			ans = append(ans, temp)
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{
		-2, 0, 0, 2, 2, 2, 4, -4, -2, 9, -6, -3,
	}))
}

/*
	总结
	1. 这题第一次错误了，因为在求2个数的和时没有去除重复的，可以看上面去重
	2. 我做出来这题的复杂度是O(n^2),我是先从左到右找一个数A，再到A数后面用twoSum找适合的元组

*/
