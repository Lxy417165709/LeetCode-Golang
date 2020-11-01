package 排序

import (
	"bytes"
	"fmt"
	"sort"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, t int) bool {
		return splice(nums[i], nums[t]) < splice(nums[t], nums[i])
	})
	return stringfy(nums)
}

func splice(A, B int) string {
	return fmt.Sprintf("%d%d", A, B)
}

func stringfy(nums []int) string {
	buffer := bytes.Buffer{}
	for _, num := range nums {
		buffer.WriteString(fmt.Sprintf("%d", num))
	}
	return buffer.String()
}

/*
	题目链接: https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
	总结
		1. 这题还可以先把数组元素转为字符串，再使用sort.Slice进行排序
		2. 其实这题有点贪心的味道 --- 只要能组成最小，就把你放最前
*/
