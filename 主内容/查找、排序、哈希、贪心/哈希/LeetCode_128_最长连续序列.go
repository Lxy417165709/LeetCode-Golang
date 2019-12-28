package main

func longestConsecutive(nums []int) int {
	isLiving := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		isLiving[nums[i]] = true
	}
	ans := 0
	for len(isLiving) != 0 {
		count := 1
		key := 0
		// 取一个键值
		for key = range isLiving {
			delete(isLiving, key)
			break
		}
		left, right := key-1, key+1
		// 向左拓展
		for isLiving[left] {
			delete(isLiving, left)
			left--
			count++
		}
		// 向右拓展
		for isLiving[right] {
			delete(isLiving, right)
			right++
			count++
		}
		ans = max(count, ans)
	}
	return ans
}

// 另外一种哈希解法
// 这种哈希解法只关心边界，
// 比如 [1 2 0 4 5 3],
// 在加入最后一个数3后，数值的变化为: length[3] = 6,length[0] = 6,length[5] = 6,而length[1,2,4]是不变的。
//
func longestConsecutive(nums []int) int {
	length := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		// 这里为了去重，如果length[nums[i]]有值了，那么就说明到目前为止，我们已经求得了
		// 包含nums[i]的连续序列长度了。(不一定是最长)(为什么? 看上面)
		if _, ok := length[nums[i]]; ok {
			continue
		}
		left, right := length[nums[i]-1], length[nums[i]+1]		// 表示对于nums[i]来说，左、右边还能拓展多长。
		curLength := left + right + 1
		length[nums[i]] = curLength
		length[nums[i]-left] = curLength
		length[nums[i]+right] = curLength
		ans = max(ans, curLength)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
	题目链接:
		https://leetcode-cn.com/problems/longest-consecutive-sequence/comments/		最长连续序列
*/
/*
	总结
	1. 这题我采用哈希的思想，思路是:
			(1) 先扫描一次nums数组，标记数字是否存在。
			(2) 之后遍历哈希表，做法是: 每次取出哈希表中的一个元素，向左向右拓展，更新最长长度。
									(拓展的时候记得要把键值对删除)
			(3) 最后就可以获得最长长度了。
	2. 这题还可以使用排序做。

*/
