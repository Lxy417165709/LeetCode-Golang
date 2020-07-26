package 位运算

func singleNumber(nums []int) []int {
	Num1XorNum2 := getXor(nums)
	diffBitOfNum1AndNum2 := getValueOfLowestBit(Num1XorNum2)
	numsHavingNum1, numsHavingNum2 := splitToTwoPartBaseOnValueOfBit(nums, diffBitOfNum1AndNum2)
	return []int{getXor(numsHavingNum1), getXor(numsHavingNum2)}
}

func getValueOfLowestBit(num int) int {
	return num & (-num)
}

func splitToTwoPartBaseOnValueOfBit(nums []int, valueOfBit int) ([]int, []int) {
	numsOfBitIsOne := make([]int, 0)
	numsOfBitIsZero := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if (nums[i] & valueOfBit) == 0 {
			numsOfBitIsZero = append(numsOfBitIsZero, nums[i])
		} else {
			numsOfBitIsOne = append(numsOfBitIsOne, nums[i])
		}
	}
	return numsOfBitIsZero, numsOfBitIsOne
}

func getXor(nums []int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}
	return xor
}

/*
	题目链接: https://leetcode-cn.com/problems/single-number-iii/comments/
*/
