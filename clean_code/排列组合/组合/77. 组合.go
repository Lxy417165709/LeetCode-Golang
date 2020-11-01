package 组合



var combinations [][]int

func combine(n int, k int) [][]int {
	combinations = make([][]int,0)
	array := make([]int,0)
	for i:=1;i<=n;i++{
		array = append(array,i)
	}
	formCombinations(array,[]int{},k)
	return combinations
}

func formCombinations(array []int,nowCombinations []int,countOfWantToSelect int) {
	if len(nowCombinations)==countOfWantToSelect{
		combinations = append(combinations,NewSlice(nowCombinations))
		return
	}
	for i:=0;i<len(array);i++{
		formCombinations(array[i+1:],append(nowCombinations,array[i]),countOfWantToSelect)
	}
}

func NewSlice(oldSlice []int) []int {
	newSlice := make([]int, len(oldSlice))
	copy(newSlice, oldSlice)
	return newSlice
}

/*
	题目链接: https://leetcode-cn.com/problems/combinations/
	题意: 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
*/
