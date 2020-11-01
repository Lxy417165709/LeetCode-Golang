package 图论

// 这个题只有一个法官
func findJudge(N int, trust [][]int) int {
	inDegree := make([]int, N+1)
	outDegree := make([]int, N+1)
	for i := 0; i < len(trust); i++ {
		outDegree[trust[i][0]]++
		inDegree[trust[i][1]]++
	}
	for i := 1; i <= N; i++ {
		if inDegree[i] == N-1 && outDegree[i] == 0 {
			return i
		}
	}
	return -1
}

/*
	总结
	1. 这题计算出入度就可以了，小镇法官的出度是0，入度是N-1
	2. 之前做的时候居然把trust构成一个邻接矩阵后再判断，真的是麻烦...
*/
