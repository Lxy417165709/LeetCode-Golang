package 深度优先搜索

import "sort"

func makesquare(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%4 != 0 || sum == 0 {
		return false
	}
	// 因为是形成4根，所以下面要传入 []int{0,0,0,0}
	return makeSequareExec(nums, []int{0, 0, 0, 0})
}

// nums: 表示原数组                          (DFS过程中不变)
// lengthSet: 表示正方形边的情况			 (DFS过程中会改变)
// 这个代码会超时
func makeSequareExec(nums []int, lengthSet []int) bool {
	if len(nums) == 0 {
		for i := 1; i < len(lengthSet); i++ {
			if lengthSet[i] != lengthSet[i-1] {
				return false
			}
		}
		return true
	}
	ans := false
	for i := 0; i < len(lengthSet); i++ {
		lengthSet[i] += nums[len(nums)-1]
		// 这其实还可以进行剪枝的，比如比较此时lengthSet[i]的长度是否大于了要形成的正方形的边长，大于就返回false
		// 但是这样的话，就要传入一个 要形成的正方形的边长 参数了。
		// 下面的函数就实现了这一步剪枝
		ans = ans || makeSequareExec(nums[:len(nums)-1], lengthSet)
		lengthSet[i] -= nums[len(nums)-1]
	}
	return ans
}

// ---------------------------------------------------- 以下是剪枝的代码 ------------------------------

// 以下是剪枝的代码
func makesquare(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%4 != 0 || sum == 0 {
		return false
	}
	sort.Ints(nums) // 这一句能极大的降低时间花费
	return makeSequareExec(nums, sum/4, []int{0, 0, 0, 0})
}

// 参数添加了targetLength字段，表示要组成的正方形的边长
func makeSequareExec(nums []int, targetLength int, lengthSet []int) bool {
	if len(nums) == 0 {
		for i := 1; i < len(lengthSet); i++ {
			if lengthSet[i] != lengthSet[i-1] {
				return false
			}
		}
		return true
	}
	ans := false
	for i := 0; i < len(lengthSet); i++ {
		// 剪枝
		if lengthSet[i]+nums[len(nums)-1] > targetLength {
			continue
		}
		lengthSet[i] += nums[len(nums)-1]
		ans = ans || makeSequareExec(nums[:len(nums)-1], targetLength, lengthSet)
		lengthSet[i] -= nums[len(nums)-1]
	}
	return ans
}

/*
	总结
	1. 这个递归代码更加简洁，而且也没使用到外部变量，我觉得很好。 (这是借鉴了大佬的代码后自己写出来的)
	2. 使用这个解法速度不太宽，但是内存消耗少了一些。   	时空花费: 368 ms 2.1 MB
		在添加了sort.Ints(nums)函数后，时间花费少了很多。	时空花费: 20 ms	2.1 MB
*/
