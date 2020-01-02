package main

func numJewelsInStones(J string, S string) int {
	jewel := make(map[uint8]bool)
	for i := 0; i < len(J); i++ {
		jewel[J[i]] = true
	}
	ans := 0
	for i := 0; i < len(S); i++ {
		if jewel[S[i]] {
			ans++
		}
	}
	return ans
}

/*
	总结
	1. 这题简单哈希就可以了。
	2. 可以使用bitmap省空间
*/
