package main



func climbStairs(n int) int {
	return getCountOfPermutation([]int{1,2},n)
}

func getCountOfPermutation(arrayOfElementGreaterThanZero []int, sumOfGreaterThanOrEqualZero int) int {
	countOfPermutationWithSpecificSum := make([]int, sumOfGreaterThanOrEqualZero+1)
	countOfPermutationWithSpecificSum[0] = 1
	for i := 1; i <= sumOfGreaterThanOrEqualZero; i++ {
		for t := 0; t < len(arrayOfElementGreaterThanZero); t++ {
			if i-arrayOfElementGreaterThanZero[t] >= 0 {
				countOfPermutationWithSpecificSum[i] += countOfPermutationWithSpecificSum[i-arrayOfElementGreaterThanZero[t]]
			}
		}
	}
	return countOfPermutationWithSpecificSum[sumOfGreaterThanOrEqualZero]
}

/*
	题目链接: https://leetcode-cn.com/problems/climbing-stairs/
*/
