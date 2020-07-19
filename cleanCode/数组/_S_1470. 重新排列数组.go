package 数组

// --------------------- 原地算法 ---------------------
func shuffle(nums []int, n int) []int {
	for i := 0; i < len(nums); i++ {
		for desireIndex := getDesireIndex(n, i); nums[desireIndex] > 0; desireIndex = getDesireIndex(n, desireIndex) {
			nums[desireIndex], nums[i] = nums[i], nums[desireIndex]
			nums[desireIndex] = -nums[desireIndex] // 负数表示该位置无需调整
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = -nums[i]
	}
	return nums
}

func getDesireIndex(n int, index int) int {
	if index < n {
		return index * 2
	}
	return 2*(index-n) + 1
}
