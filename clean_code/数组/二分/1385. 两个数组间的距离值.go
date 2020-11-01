package 二分

import "sort"

// ---------------------------- 方法1: 暴力版 ----------------------------
const INF = 1000000000

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	return getDistanceValueOfFirstArrayToSecondArray(arr1, arr2, d)

}

func getDistanceValueOfFirstArrayToSecondArray(firstArr, secondArr []int, d int) int {
	distanceValue := 0
	for i := 0; i < len(firstArr); i++ {
		if getMinDistance(secondArr, firstArr[i]) > d {
			distanceValue++
		}
	}
	return distanceValue
}

func getMinDistance(arr []int, ref int) int {
	minDistance := INF
	for i := 0; i < len(arr); i++ {
		minDistance = min(minDistance, abs(arr[i]-ref))
	}
	return minDistance
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// ---------------------------- 方法2: 二分查找版 ----------------------------
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	return getDistanceValueOfFirstArrayToSecondArray(arr1, arr2, d)
}

func getDistanceValueOfFirstArrayToSecondArray(firstArr, secondArr []int, d int) int {
	distanceValue := 0
	sort.Ints(secondArr)
	for i := 0; i < len(firstArr); i++ {
		indexOfLastLessOrEqual := getIndexOfLastLessOrEqual(secondArr, firstArr[i])
		if indexOfLastLessOrEqual != -1 && getDistance(firstArr[i], secondArr[indexOfLastLessOrEqual]) <= d {
			continue
		}
		indexOfFirstGreaterOrEqual := getIndexOfFirstGreaterOrEqual(secondArr, firstArr[i])
		if indexOfFirstGreaterOrEqual != len(secondArr) && getDistance(firstArr[i], secondArr[indexOfFirstGreaterOrEqual]) <= d {
			continue
		}
		distanceValue++
	}
	return distanceValue
}

func getDistance(a, b int) int {
	return abs(a - b)
}

func getIndexOfLastLessOrEqual(arr []int, ref int) int {
	return getIndexOfFirstGreater(arr, ref) - 1
}

func getIndexOfFirstGreaterOrEqual(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func getIndexOfFirstGreater(arr []int, ref int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}


/*
	题目链接: https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays/submissions/
*/
