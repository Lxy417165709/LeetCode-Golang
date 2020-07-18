package 二分

import "math/rand"

func smallestK(arr []int, k int) []int {
	makeKthSmallToRightPlace(arr,k)
	return arr[:k]
}


func makeKthSmallToRightPlace(nums []int, KthSmall int) {
	l, r := 0, len(nums)-1
	for l <= r {
		index := randomPartition(nums, l, r)
		XthSmall := index+1
		if XthSmall == KthSmall {
			return
		}

		if XthSmall > KthSmall {
			r = index - 1
		} else {
			l = index + 1
		}

	}
}

func makeKthBigToRightPlace(nums []int, KthBig int) {
	makeKthSmallToRightPlace(nums, len(nums)-KthBig+1)
}

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
	题目链接: https://leetcode-cn.com/problems/smallest-k-lcci/submissions/
*/
