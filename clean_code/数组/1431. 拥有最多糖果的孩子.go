package 数组

const INF = 1000000000000

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := getMax(candies)
	canBeMostCandiesKid := make([]bool, 0)
	for i := 0; i < len(candies); i++ {
		canBeMostCandiesKid = append(canBeMostCandiesKid, candies[i]+extraCandies >= maxCandies)
	}
	return canBeMostCandiesKid
}

func getMax(nums []int) int {
	result := -INF
	for i := 0; i < len(nums); i++ {
		result = max(nums[i], result)
	}
	return result
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a := arr[0]
	b := max(arr[1:]...)
	if a < b {
		return b
	}
	return a
}

/*
	题目链接: https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/submissions/
*/
