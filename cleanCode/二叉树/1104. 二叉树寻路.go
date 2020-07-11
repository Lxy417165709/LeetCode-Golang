package 二叉树

func pathInZigZagTree(label int) []int {
	labels := make([]int, 0)
	for label != -1 {
		labels = append(labels, label)
		label = 3*getValueOfHighestOne(label/2)-label/2-1
	}
	return reverse(labels)
}


func getValueOfHighestOne(num int) int {
	valueOfHighestOne := 0
	for num != 0 {
		valueOfLowBit := getValueOfLowestBit(num)
		valueOfHighestOne = max(valueOfLowBit,valueOfHighestOne)
		num ^= valueOfLowBit
	}
	return valueOfHighestOne
}

func getValueOfLowestBit(num int) int {
	return num & (-num)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(nums []int) []int{
	length := len(nums)
	for i:=0;i<length/2;i++{
		nums[i],nums[length-1-i] = nums[length-1-i],nums[i]
	}
	return nums
}


/*
	题目链接: https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/
	总结:
		1. 感觉这题就是找规律。
*/
