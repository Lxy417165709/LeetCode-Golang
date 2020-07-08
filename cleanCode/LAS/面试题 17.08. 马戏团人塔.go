package main

func bestSeqAtIndex(heights []int, weights []int) int {
	persons := getPersons(heights, weights)
	sortPersons(persons)
	return getLengthOfLAS(getWeightsOfAllPersons(persons))
}

func getPersons(heights []int, weights []int) []*Person {
	persons := make([]*Person, 0)
	for i := 0; i < len(heights); i++ {
		persons = append(persons, NewPerson(heights[i], weights[i]))
	}
	return persons
}

func sortPersons(persons []*Person) {
	sort.Slice(persons, func(i, t int) bool {
		if persons[i].Height != persons[t].Height {
			return persons[i].Height < persons[t].Height
		}
		return persons[i].Weight > persons[t].Weight
	})
}

func getWeightsOfAllPersons(persons []*Person) []int {
	weights := make([]int, 0)
	for i := 0; i < len(persons); i++ {
		weights = append(weights, persons[i].Weight)
	}
	return weights
}

func getLengthOfLAS(nums []int) int {
	LAS := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := getIndexOfFirstGreaterOrEqual(LAS, nums[i])
		if index == len(LAS) {
			LAS = append(LAS, nums[i])
		}
		LAS[index] = nums[i]
	}
	return len(LAS)
}

func getIndexOfFirstGreaterOrEqual(nums []int, ref int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= ref {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// --------------- Person -----------------
type Person struct {
	Height int
	Weight int
}

func NewPerson(height, weight int) *Person {
	return &Person{height, weight}
}

/*
	题目链接: https://leetcode-cn.com/problems/circus-tower-lcci/
	总结
		1. 这题刚刚开始排序规则定错了，所以WA，看了题解后，发现按高度升序，宽度降序就能AC了。
				(这种排序规则就能避免高度相同时，宽度还能进入LAS中)
		2. 这题结合了 排序 + LAS + 二分
		3. 为了提高可复用性，我把LAS抽象了出来
*/
