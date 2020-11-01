package 单数

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{k * longer}
	}
	divingBoards := make([]int, 0)
	for countOfLonger := 0; countOfLonger <= k; countOfLonger++ {
		countOfShorter := k - countOfLonger
		length := countOfLonger*longer + countOfShorter*shorter
		divingBoards = append(divingBoards, length)
	}
	return divingBoards
}
