package main

func getLeastNumbers(arr []int, k int) []int {
	makeKthSmallToRightPlace(arr,k)
	return arr[:k]
}

// 例子: A 的 rightPlace 就是：A处于某个位置，这个位置左边的元素都小于等于它，右边的元素都大于等于它。
func makeKthSmallToRightPlace(nums []int, KthSmall int) {
	l, r := 0, len(nums)-1
	for l <= r {
		index := randomPartition(nums, l, r)
		XthSmall := index+1
		if XthSmall == KthSmall {
			return
		} else {
			if XthSmall > KthSmall {
				r = index - 1
			} else {
				l = index + 1
			}
		}
	}
	// 到达这里表示:  原数组不存在第K小
}

func makeKthBigToRightPlace(nums []int, KthBig int) {
	makeKthSmallToRightPlace(nums, len(nums)-KthBig+1)
}

// 返回值是一个索引，假如记做 index
// 则 nums[index] 位于 rightPlace
// 同 partition函数
func randomPartition(nums []int, l int, r int) int {
	upsetArrayRandomly(nums,l,r)
	return partition(nums,l,r)
}

func upsetArrayRandomly(nums []int, l int, r int) {
	randomIndex := rand.Intn(r-l+1) + l
	nums[randomIndex], nums[l] = nums[l], nums[randomIndex]
}

func partition(nums []int, l int, r int) int {
	guardIndex := l
	for l <= r {
		for l <= r && nums[l] <= nums[guardIndex] {
			l++
		}
		for l <= r && nums[r] >= nums[guardIndex] {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[guardIndex], nums[r] = nums[r], nums[guardIndex]
	return r
}






/*

	题目链接:
	总结
		1. 这题我是使用随机选择算法AC的，题解还有用排序、堆解决的。
		2. l, r 可以抽象为区间
		3. l其实可以命名为left,这样 l 就不会像 数字1 了。
*/
