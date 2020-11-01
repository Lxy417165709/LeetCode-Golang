package 深度优先搜索

func makesquare(nums []int) bool {
	isUsed = 0
	isVisit = make(map[int]bool)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%4 != 0 || sum == 0 {
		return false
	}
	length := sum / 4
	return makeSequareExec(nums, length, 0, 0)
}

var isUsed int // 利用位运算来实现map，isUsed的二进制每一位表示的是使用情况
// 假如 isUsed == 6，6的二进制为 110 (从右到左为第0、第1、第2...)，
// 那么表示nums[1],nums[2]已经使用了。

var isVisit map[int]bool // 一个备忘录 (键值是 数组的元素的使用情况(对应isUsed变量) )

// nums: 表示原数组                         (DFS过程中不变)
// targetLength: 表示要形成的正方形的边长    (DFS过程中不变)
// nowLength: 表示到目前为止已经构成的长度   (DFS过程中会变)
// count: 表示此时已经形成的targetLength长度的火柴条个数。(火柴可以是由多根火柴拼接而成) (DFS过程中会变)
func makeSequareExec(nums []int, targetLength int, nowLength int, count int) bool {
	if x, ok := isVisit[isUsed]; ok {
		return x
	}
	if count == 4 {
		// 达到4根了，那么就要判断一下数组是否还有剩余火柴
		return isUsed == (1<<len(nums))-1
	}
	// 剪枝
	if nowLength > targetLength {
		return false
	}
	if nowLength == targetLength {
		return makeSequareExec(nums, targetLength, 0, count+1)
	}
	ans := false
	for i := 0; i < len(nums); i++ {
		if isUsed&(1<<i) != 0 {
			continue
		}
		isUsed = isUsed ^ (1 << i)
		ans = ans || makeSequareExec(nums, targetLength, nowLength+nums[i], count)
		isUsed = isUsed ^ (1 << i)
	}
	isVisit[isUsed] = ans
	return ans
}

/*
	总结
	1. 这题采用了位运算 + 备忘录 + 回溯。
			思路是:
				(1) 先获取正方形的边长
				(2) 通过搜索，判断数组能否正好形成4根等于边长的火柴(火柴可以是拼接而成的)
	2. 使用这个解法速度很快，但是内存消耗多了一些。   时空花费: 16 ms 3.8 MB
	3. 官方有其他的递归思路喔。
*/
