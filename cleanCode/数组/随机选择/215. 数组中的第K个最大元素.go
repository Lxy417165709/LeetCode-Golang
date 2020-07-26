package 随机选择

const kthNotExist = -10000000000000

func findKthLargest(nums []int, k int) int {
	return getKthBig(nums, k)
}

func getKthBig(nums []int, k int) int {
	return getKthSmall(nums, len(nums)-k+1)
}

// 要保证 KthSmall 存在。
func getKthSmall(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		order := partition(nums, left, right) + 1
		if order == k {
			return nums[order-1]
		}
		if order > k {
			right = order - 2
		} else {
			left = order
		}
	}
}

func partition(nums []int, left, right int) int {
	guardIndex := left
	guard := nums[guardIndex]
	for left <= right {
		if left <= right && nums[left] <= guard {
			left++
		}
		if left <= right && nums[right] >= guard {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[right], nums[guardIndex] = nums[guardIndex], nums[right]
	return right
}
