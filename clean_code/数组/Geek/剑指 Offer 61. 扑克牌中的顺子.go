package Geek

import "sort"

const sizeOfStraight = 5

func isStraight(nums []int) bool {
	sort.Ints(nums)
	tail, head := getHeadAndTailOfStraight(nums)
	return tail-head < sizeOfStraight && !hasRepeatNum(nums)
}

func hasRepeatNum(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 && nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func getHeadAndTailOfStraight(nums []int) (int, int) {
	i := 0
	for nums[i] == 0 {
		i++
		if i == len(nums) {
			return 0, 0
		}
	}
	return nums[len(nums)-1], nums[i]
}
