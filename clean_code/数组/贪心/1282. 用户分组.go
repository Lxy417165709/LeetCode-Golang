package 贪心

func groupThePeople(groupSizes []int) [][]int {
	uids := make(map[int][]int)
	result := make([][]int, 0)
	for i := 0; i < len(groupSizes); i++ {
		uid := i
		uids[groupSizes[i]] = append(uids[groupSizes[i]], uid)
		if len(uids[groupSizes[i]]) == groupSizes[i] {
			result = append(result, uids[groupSizes[i]])
			delete(uids, groupSizes[i])
		}
	}
	return result
}

/*
	题目链接: https://leetcode-cn.com/problems/group-the-people-given-the-group-size-they-belong-to/
*/
