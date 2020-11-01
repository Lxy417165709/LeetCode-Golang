package main

var sequence [][]int

func allPathsSourceTarget(graph [][]int) [][]int {
	sequence = make([][]int, 0)
	allPathsSourceTargetExec(graph, []int{0})
	return sequence
}

func allPathsSourceTargetExec(graph [][]int, nowSequence []int) {
	nowPoint := nowSequence[len(nowSequence)-1]
	if nowPoint == len(graph)-1 {
		sequence = append(sequence, newSlice(nowSequence))
		return
	}
	for i := 0; i < len(graph[nowPoint]); i++ {
		allPathsSourceTargetExec(graph, append(nowSequence, graph[nowPoint][i]))
	}
}

func newSlice(oldSlice []int) []int {
	s := make([]int, len(oldSlice))
	copy(s, oldSlice)
	return s
}

/*
	总结
	1. 这题就是对邻接表进行遍历，得出所有从源点到达终点的路径。
*/
