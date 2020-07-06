package 前缀和___哈希


func findMaxLength(nums []int) int {
	indexOfDiffCountOfOneAndZero := make(map[int]int)
	diffCountOfOneAndZero := 0
	indexOfDiffCountOfOneAndZero[0] = -1
	maxLength := 0
	for i:=0;i<len(nums);i++{
		if nums[i]==1{
			diffCountOfOneAndZero++
		}else{
			diffCountOfOneAndZero--
		}
		if index,ok := indexOfDiffCountOfOneAndZero[diffCountOfOneAndZero];!ok{
			indexOfDiffCountOfOneAndZero[diffCountOfOneAndZero]=i
		}else{
			maxLength = max(maxLength,i-index)
		}
	}
	return maxLength
}


func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/contiguous-array/submissions/
	总结:
		1. 解题思路和 面试题 17.05.  字母与数字 一样。
		2. 其实这题也有点贪心的思维 --- 只记录第一次出现的差值与索引的映射关系
		3. 个人感觉: 抽象成数学模型就是，找出2个值一样的点的最大距离。
*/
