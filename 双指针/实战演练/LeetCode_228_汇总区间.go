package main

import "fmt"

func summaryRanges(nums []int) []string {
	ans := []string{}
	if len(nums) == 0 {
		return []string{}
	}

	begin, end := 0, 1 // [begin,end) 区间内为连续的数字
	for end < len(nums) {
		if nums[end] == nums[end-1]+1 {
			end++
			continue
		}
		ans = append(ans, toString(nums, begin, end))
		begin = end
		end++
	}
	ans = append(ans, toString(nums, begin, end))
	return ans
}

// 这只是为了满足题目的格式要求

// 代码更简化
func summaryRanges(nums []int) []string {
	ans := []string{}
	begin, end := 0, 1
	for end <= len(nums) {
		if end == len(nums) || nums[end] != nums[end-1]+1 {
			ans = append(ans, toString(nums, begin, end))
			begin = end
		}
		end++
	}
	return ans
}

func toString(nums []int, begin, end int) string {
	str := ""
	if end-begin == 1 {
		str = fmt.Sprintf("%d", nums[begin])
	} else {
		str = fmt.Sprintf("%d->%d", nums[begin], nums[end-1])
	}
	return str
}

/*
	题目链接:
		https://leetcode-cn.com/problems/summary-ranges/submissions/	 汇总区间
*/
/*
	总结
	1. 这道题我使用了双指针。
	2. 在循环体外，我还 ans = append(ans, toString(nums, begin, end))，这是因为end==len(nums)时，此时[begin,end)区间我们
       还没有加入结果集。究其本质就是: nums[-1]与nums[-2...]有关系。	(对于解法1)
	3. 对于解法2，我直接让end循环多一次，当end==len(nums)时，也直接加入结果集，这样就不用在循环体外操作了。
*/
