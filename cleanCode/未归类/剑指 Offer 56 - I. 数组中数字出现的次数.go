package 未归类

func singleNumbers(nums []int) []int {
	AxorB := getXor(nums)
	lowestBitValueOfAxorB := getLowestBitValue(AxorB)
	numsExistingA, numsExistingB := splitToTwoPilesByBitValue(nums, lowestBitValueOfAxorB)
	A := getXor(numsExistingA)
	B := getXor(numsExistingB)
	return []int{A, B}
}

func getXor(nums []int) int {
	Xor := 0
	for _, num := range nums {
		Xor = Xor ^ num
	}
	return Xor
}

func getLowestBitValue(num int) int {
	return num & (-num)
}

func splitToTwoPilesByBitValue(nums []int, bitValue int) ([]int, []int) {
	numsOfPositionOfBitIsZero := make([]int, 0)
	numsOfPositionOfBitIsNotZero := make([]int, 0)
	for _, num := range nums {
		if isPositionOfBitZero(num, bitValue) {
			numsOfPositionOfBitIsZero = append(numsOfPositionOfBitIsZero, num)
		} else {
			numsOfPositionOfBitIsNotZero = append(numsOfPositionOfBitIsNotZero, num)
		}
	}
	return numsOfPositionOfBitIsZero, numsOfPositionOfBitIsNotZero
}

func isPositionOfBitZero(num int, bitValue int) bool {
	return (num & bitValue) == 0
}

/*
	链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
*/
