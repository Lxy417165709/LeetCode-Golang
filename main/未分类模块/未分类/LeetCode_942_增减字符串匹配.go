package main

func diStringMatch(S string) []int {
	ans := []int{}
	mn := 0
	mx := len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 'D' {
			ans = append(ans, mx)
			mx--
		} else {
			ans = append(ans, mn)
			mn++
		}
	}
	ans = append(ans, mn)
	return ans
}

/*
	总结
	1. 这题类似贪心的思维，当遇到'D'时，表明要下降了，那么为了保证下个数一定下降，
		我们就取区间的最大值填入，而遇到'I'，表明要上升了，那么为了保证下个数一定上升，
		我们就区区间最小值填入。
*/
