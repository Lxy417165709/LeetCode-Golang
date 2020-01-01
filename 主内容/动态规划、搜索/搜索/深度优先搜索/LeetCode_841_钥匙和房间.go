package 深度优先搜索

var isVisit map[int]bool

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	isVisit = make(map[int]bool)
	canVisitAllRoomsExec(rooms, 0)
	return len(rooms) == len(isVisit)
}

func canVisitAllRoomsExec(rooms [][]int, nowVisit int) {
	if isVisit[nowVisit] == true || len(rooms) <= nowVisit {
		return
	}
	isVisit[nowVisit] = true
	for i := 0; i < len(rooms[nowVisit]); i++ {
		canVisitAllRoomsExec(rooms, rooms[nowVisit][i])
	}
}

/*
	总结
	1. 这题目简单DFS就可以了，标记每个房间房间是否访问，最后如果都访问就true，否则false。
	2. 上面采用了map，可能速度不是很快，接下来采用数组。
*/
