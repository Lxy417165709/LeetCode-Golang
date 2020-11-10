package 二分典型题

const inf = 1000000000000

func splitArray(nums []int, m int) int {
	left, right := 0, inf
	for left <= right {
		mid := (left + right) / 2
		countOfSegment := getCountOfSegment(nums, mid)
		if countOfSegment == m {
			right = mid - 1
		} else if countOfSegment > m {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func getCountOfSegment(nums []int, maxSum int) int {
	countOfSegment := 0
	curSum := 0
	for _, num := range nums {
		if num > maxSum {
			return inf
		}
		if curSum+num > maxSum {
			curSum = num
			countOfSegment++
			continue
		}
		curSum = curSum + num
	}
	if curSum != 0 {
		countOfSegment++
	}
	return countOfSegment
}
