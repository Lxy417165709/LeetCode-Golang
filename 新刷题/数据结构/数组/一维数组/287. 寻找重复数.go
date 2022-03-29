package 一维数组

import "fmt"

// findDuplicate 寻找数组中重复的数字。 (二分法)
func findDuplicate(nums []int) int {
	// 1. 定义数值边界。
	left, right := 1, len(nums)-1

	// 2. 二分。
	for left <= right {
		// 2.1 计算中间数值在数组中出现的次数。
		mid := (left + right) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		// 2.2 边界收缩。
		if mid == count {
			left = mid + 1
		} else if mid > count {
			left = mid + 1
		} else if mid < count {
			right = mid - 1 // 说明 [1, mid] 中至少有 [mid+1]个数，此时重复元素一定在 [1, mid] 中。
		}
	}

	// 3. 返回。
	return left
}

// findDuplicate 寻找数组中重复的数字。 (快慢指针法)
// 链表的构成:
// 1. 节点是数组索引。
// 2. 边是数组数字。
// 例子:
// nums = [1, 2, 3, 1]
// 此时构成的链表为
//                 ----------------------
//                 |                     |
//                 1                     |
//                 ↓                     |
//     [0] --1--> [1] --2--> [2] --3--> [3]
//
// 证明:
// 1. 为什么链表一定存在环？因为链表有 n 个节点、n 条边，所以一定有环。
func findDuplicate(nums []int) int {
	// 1. 快慢指针入环。
	slow, fast := 0, 0 // 从链表起始点出发。
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// 2. 确定环口。
	cur := 0
	for cur != slow {
		cur = nums[cur]
		slow = nums[slow]
	}

	// 3. 返回。
	return slow
}
