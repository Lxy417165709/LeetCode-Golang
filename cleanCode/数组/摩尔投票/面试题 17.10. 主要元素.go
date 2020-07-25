package 摩尔投票

func majorityElement(nums []int) int {
	candidateNum := getMajorElementCandidate(nums)
	minCountOfMajorElement := len(nums)/2 + 1
	if getCountOfNum(nums, candidateNum) >= minCountOfMajorElement {
		return candidateNum
	} else {
		return -1
	}
}
func getMajorElementCandidate(nums []int) int {
	countOfCandidateNum := 0
	candidateNum := 0
	for i := 0; i < len(nums); i++ {
		if countOfCandidateNum == 0 {
			countOfCandidateNum = 1
			candidateNum = nums[i]
			continue
		}
		if candidateNum == nums[i] {
			countOfCandidateNum++
		} else {
			countOfCandidateNum--
		}
	}
	return candidateNum
}

func getCountOfNum(arr []int, num int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			count++
		}
	}
	return count
}

/*
	总结:
		1. 摩尔投票假定数组中一定存在主要元素，所以这题在得到"主要元素候选者"后，需要对
			该候选者进行校验，判断它是不是主要元素。
*/
