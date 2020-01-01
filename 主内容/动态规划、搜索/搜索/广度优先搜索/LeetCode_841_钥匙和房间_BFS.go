package 广度优先搜索

func canVisitAllRooms(rooms [][]int) bool {
	return BFS(rooms, []int{0})
}

func BFS(rooms [][]int, queue []int) bool {
	if len(rooms) == 0 {
		return true
	}
	isVisit := make([]bool, len(rooms))
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top >= len(isVisit) || isVisit[top] == true {
			continue
		}
		isVisit[top] = true
		for i := 0; i < len(rooms[top]); i++ {
			queue = append(queue, rooms[top][i])
		}
	}
	for i := 0; i < len(isVisit); i++ {
		if isVisit[i] == false {
			return false
		}
	}
	return true
}

/*
	总结
	1. 这题目简单DFS就可以了，标记每个房间房间是否访问，最后如果都访问就true，否则false。
	2. 上面采用了map，可能速度不是很快，接下来采用数组。
	3. 如果有bitmap，其实最好用bitmap，这样很省空间和时间。
	4. 这题还可以使用队列实现。
	5. 以上就是采用BFS实现的代码
*/
