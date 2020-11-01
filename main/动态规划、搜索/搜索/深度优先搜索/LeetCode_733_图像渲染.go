package 深度优先搜索

var dx, dy []int
// 判断坐标需要渲染
func judge(image [][]int, x, y int, baseColor int) bool {
	if x < 0 || y < 0 || x >= len(image) || y >= len(image[0]) {
		return false
	}
	return image[x][y] == baseColor
}

// 这就是个DFS
func floodFillExec(image [][]int, x, y int, baseColor int, newColor int) {
	if judge(image, x, y, baseColor) == false {
		return
	}
	image[x][y] = newColor
	for i := 0; i < len(dx); i++ {
		floodFillExec(image, x+dx[i], y+dy[i], baseColor, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dx = []int{0, 0, 1, -1}
	dy = []int{1, -1, 0, 0}
	// 判断越界
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) {
		return image
	}
	// 如果原图像点颜色等于新的图像点的颜色，那么就不用渲染了
	if image[sr][sc] == newColor {
		return image
	}
	floodFillExec(image, sr, sc, image[sr][sc], newColor)
	return image
}

/*
	总结
	1. 这题就是一道简单的DFS连通图题。就是有个细节要注意 ---> 原图像点颜色等于新的图像点的颜色时要直接返回。
*/
