package 单调栈

// 利用哈希和单调栈解决该题
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	lastBig := make(map[int]int)
	stack := make([]int, 0)
	index := len(nums2) - 1
	for index >= 0 {
		for len(stack) != 0 && nums2[index] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			lastBig[nums2[index]] = -1
		} else {
			lastBig[nums2[index]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[index])
		index--
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = lastBig[nums1[i]]
	}
	return nums1
}


/*
	题目链接:
		https://leetcode-cn.com/problems/next-greater-element-i/			下一个更大元素 I
*/

/*
	总结
	1. 这题我使用了哈希和单调递减栈进行解决。思路是: 首先逆序扫描一遍nums2，通过单调栈获取nums2[index]下一个更大的数,
	   并记录在哈希表中，之后再扫描一次nums1，把哈希结果存放在nums1中。
*/
