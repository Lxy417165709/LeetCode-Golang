package 数组

// ---------------------- getCountOfNum ------------------------
func intersection(nums1 []int, nums2 []int) []int {
	countOfNums1 := getCountOfNum(nums1)
	countOfNums2 := getCountOfNum(nums2)
	intersection:= make([]int,0)
	for num := range countOfNums1{
		if countOfNums2[num]!=0{
			intersection = append(intersection,num)
		}
	}
	return intersection
}

func getCountOfNum(array []int) map[int]int {
	count := make(map[int]int)
	for i := 0; i < len(array); i++ {
		count[array[i]]++
	}
	return count
}

// ---------------------- getIsExist ------------------------
func intersection(nums1 []int, nums2 []int) []int {
	isExistInNums1 := getIsExist(nums1)
	isExistInNums2 := getIsExist(nums2)
	intersection:= make([]int,0)
	for num := range isExistInNums1{
		if isExistInNums2[num]{
			intersection = append(intersection,num)
		}
	}
	return intersection
}

func getIsExist(array []int) map[int]bool {
	isExist := make(map[int]bool)
	for i := 0; i < len(array); i++ {
		isExist[array[i]]=true
	}
	return isExist
}


/*
	题目链接: https://leetcode-cn.com/problems/intersection-of-two-arrays/
*/
