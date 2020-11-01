package 数组

// ------------------------- 暴力法 -------------------------
// 执行用时：24 ms, 在所有 Go 提交中击败了19.48% 的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了100.00% 的用户
//
// 时间复杂度：O(n^3)
func numTeams(rating []int) int {
	countOfFightingUnit := 0
	for i := 0; i < len(rating); i++ {
		for t := i + 1; t < len(rating); t++ {
			for j := t + 1; j < len(rating); j++ {
				if rating[i] < rating[t] && rating[t] < rating[j] {
					countOfFightingUnit++
				}
				if rating[i] > rating[t] && rating[t] > rating[j] {
					countOfFightingUnit++
				}
			}
		}
	}
	return countOfFightingUnit
}

// ------------------------- 固定中间索引，两边扩展 -------------------------
// 执行用时：0 ms,   在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度：O(n^2)
func numTeams(rating []int) int {
	countOfFightingUnit := 0
	for mid := 0;mid<len(rating);mid++{
		countOfFightingUnit+=getCountOfGreaterNum(rating[:mid],rating[mid])*getCountOfLessNum(rating[mid+1:],rating[mid])
		countOfFightingUnit+=getCountOfLessNum(rating[:mid],rating[mid])*getCountOfGreaterNum(rating[mid+1:],rating[mid])
	}
	return countOfFightingUnit
}

func getCountOfGreaterNum(arr []int,ref int) int{
	count := 0
	for i:=0;i<len(arr);i++{
		if arr[i]>ref{
			count++
		}
	}
	return count
}

func getCountOfLessNum(arr []int,ref int) int{
	count := 0
	for i:=0;i<len(arr);i++{
		if arr[i]<ref{
			count++
		}
	}
	return count
}

/*
	题目链接: https://leetcode-cn.com/problems/count-number-of-teams/solution/tong-ji-zuo-zhan-dan-wei-shu-by-leetcode-solution/
	总结:
		1. 这题应该不能用单调栈做，因为它要求的元素没有连续。
*/
