package main

var father []int

func initUFS(size int) {
	father = make([]int, size)
	for i := 0; i < size; i++ {
		father[i] = i
	}
}

func mergeUFS(a int, b int) {
	father1, father2 := findFather(a), findFather(b)
	father[father1] = father2
}
func findFather(a int) int {
	if a == father[a] {
		return a
	}
	father[a] = findFather(father[a])
	return father[a]
}

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	N := len(M)
	initUFS(N)
	for i := 0; i < N; i++ {
		for t := 0; t <= i; t++ {
			if M[i][t] == 1 {
				mergeUFS(i, t)
			}
		}
	}
	/*
	hasBeenCounted := map[int]bool{}	// 这个map其实可以不用
	ans := 0
	for i := 0; i < N; i++ {
		if hasBeenCounted[findFather(i)] == true {
			continue
		}
		hasBeenCounted[findFather(i)] = true
		ans++
	}
	*/

	// 采用这种写法也可以计算连通块数
	ans := 0
	for i := 0; i < N; i++ {
		if findFather(i) == i {
			ans++
		}
	}
	return ans
}

/*
	总结
	1. 这题本质上就是求连通块的数量，所以其实还可以用BFS、DFS做。
	2. 以上我采用了并查集做。
*/
