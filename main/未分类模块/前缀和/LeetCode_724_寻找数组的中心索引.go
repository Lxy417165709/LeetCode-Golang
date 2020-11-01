package 前缀和

func pivotIndex(nums []int) int {
	sum := make([]int, len(nums)+1)
	// 获取前缀和， sum[i] 表示 nums[:i+1] 的和
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for i := 1; i <= len(nums); i++ {
		left := sum[i-1]                 // nums[:i]的和
		right := sum[len(nums)] - sum[i] // nums[i+1:]的和
		if left == right {
			return i - 1
		}
	}
	return -1

}
/*
	总结
	1. 为了方便处理，sum的下标尽量选择从1开始。
*/
