package 正确位

const flagOfMissing = -1

func missingTwo(nums []int) []int {
	nums = append(nums, flagOfMissing, flagOfMissing)
	for i := 0; i < len(nums); i++ {
		for !isRightPosition(nums, i) {
			toRightPosition(nums, i)
		}
	}
	missNum := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == flagOfMissing {
			missNum = append(missNum, i+1)
		}
	}
	return missNum
}

func isRightPosition(nums []int, index int) bool {
	return nums[index] == -1 || nums[index] == index+1
}

func toRightPosition(nums []int, index int) {
	rightPosition := getRightPosition(nums, index)
	nums[index], nums[rightPosition] = nums[rightPosition], nums[index]
}

func getRightPosition(nums []int, index int) int {
	return nums[index] - 1
}

/*
	总结:
		1. 官方还有人用位运算做这个题。
		2. 还有人用数学法解二元一次方程组的方式做出这题。
*/
