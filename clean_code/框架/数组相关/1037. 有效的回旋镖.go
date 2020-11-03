package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 将给定数组排序
 * @param arr int整型一维数组 待排序的数组
 * @return int整型一维数组
 */
func MySort( arr []int ) []int {
	// write code here
	comparableElements := make([]Comparable,0)
	for _,num := range arr{
		comparableElements = append(comparableElements,Int(num))
	}
	QuickSort(comparableElements)
	for index,element := range comparableElements{
		arr[index] = int(element.(Int))
	}
	return arr
}



type Int int

func (i Int) Greater(c Comparable) bool {
	j := c.(Int)
	return i > j
}


type Comparable interface {
	Greater(c Comparable) bool
}

// ---------------- 排序相关 ----------------
func QuickSort(nums []Comparable) {
	if len(nums) == 0 {
		return
	}
	rightPositionIndex := Partition(nums)
	QuickSort(nums[:rightPositionIndex])
	QuickSort(nums[rightPositionIndex+1:])
}

func Partition(nums []Comparable) int {
	refNumIndex := 0
	left, right := 0, len(nums)-1
	for left <= right {
		for left <= right && isSmallerOrEqual(nums[left], nums[refNumIndex]) {
			left++
		}
		for left <= right && isGreaterOrEqual(nums[right], nums[refNumIndex]) {
			right--
		}
		if left <= right {
			nums[right], nums[left] = nums[left], nums[right]
		}
	}
	nums[right], nums[refNumIndex] = nums[refNumIndex], nums[right]
	return right
}

func isEqual(firstNum Comparable, secondNum Comparable) bool {
	return !firstNum.Greater(secondNum) && !secondNum.Greater(firstNum)
}

func isGreaterOrEqual(firstNum Comparable, secondNum Comparable) bool {
	return firstNum.Greater(secondNum) || isEqual(firstNum, secondNum)
}

func isSmallerOrEqual(firstNum Comparable, secondNum Comparable) bool {
	return !firstNum.Greater(secondNum)
}


/*
	题目链接: https://www.nowcoder.com/practice/2baf799ea0594abd974d37139de27896?tpId=188&&tqId=36195&rp=1&ru=/ta/job-code-high-week&qru=/ta/job-code-high-week/question-ranking
*/