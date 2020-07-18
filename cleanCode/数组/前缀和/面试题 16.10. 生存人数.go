package 前缀和

const minYear, maxYear = 1900, 2000

func maxAliveYear(birth []int, death []int) int {
	countOfPersonalChange := getCountOfPersonChange(birth, death)
	countOfPerson := getPrefixSum(countOfPersonalChange)
	return getIndexOfMaxValue(countOfPerson) + minYear
}

func getCountOfPersonChange(birth []int, death []int) []int {
	countOfPersonalChange := make([]int, maxYear-minYear+2)
	for i := 0; i < len(birth); i++ {
		countOfPersonalChange[birth[i]-minYear]++
		countOfPersonalChange[death[i]-minYear+1]--
	}
	return countOfPersonalChange
}

func getPrefixSum(arrayOfRemainingChange []int) []int {
	prefixSum := make([]int, len(arrayOfRemainingChange))
	prefixSum[0] = arrayOfRemainingChange[0]
	for i := 1; i < len(arrayOfRemainingChange); i++ {
		prefixSum[i] = prefixSum[i-1] + arrayOfRemainingChange[i]
	}
	return prefixSum
}

func getIndexOfMaxValue(arr []int) int {
	indexOfMaxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[indexOfMaxValue] {
			indexOfMaxValue = i
		}
	}
	return indexOfMaxValue
}

/*
	题目链接: https://leetcode-cn.com/problems/living-people-lcci/
	总结:
		1. 记录变化是一个很好的编程手段。
*/
