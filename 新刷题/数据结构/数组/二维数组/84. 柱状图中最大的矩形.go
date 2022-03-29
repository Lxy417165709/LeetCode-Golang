package 二维数组

// largestRectangleArea 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1，求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 前提提示:
// 1. 最大面积时，一定有一个板子被完全包含。 (以下称为短板) (这个可以用反证法证明)
//
// 暴力:
// 1. 把每个板子当做短板，向左、右两侧扩展，直到遇到更短的板子，向左右扩展的长度 +1 为长方形宽度，短板长度即为长方形高度。
//
// 单调栈: (一定得是单调非递减栈，场景: nums[2 2])
// 1. 单调非递减栈: 栈顶元素即为短板，栈顶元素与第二个元素之间的距离为左扩展长度，与即将加入栈顶的元素(假如比栈顶小)的距离即为右扩展长度。
// 2. 栈内可能是: [2 2 3] (栈顶在右边)，当加入板子1时，此时要注意3是短板，不是长板。
func largestRectangleArea(heights []int) int {
	// 1. 初始化。
	indexStack := make([]int, 0) // 单调递增栈，存放元素索引。

	// 2. 添加无穷小高度。 (更好的方式是左右两侧都加0)
	heights = append(heights, 0) // 添加最小高度 0，保证 for 循环后，单调非递减栈内的元素只剩 0，保证所有短板都能被处理。

	// 3. 处理。
	result := 0
	for i := 0; i < len(heights); i++ {
		// 3.1 维护单调非递减特性。
		for len(indexStack) != 0 && heights[i] < heights[indexStack[len(indexStack)-1]] {
			// 3.1.1 获取矩形高度。
			shortBoardIndex := indexStack[len(indexStack)-1]
			matrixHeight := heights[shortBoardIndex]

			// 3.1.2 出栈。
			indexStack = indexStack[:len(indexStack)-1]

			// 3.1.3 获取矩阵宽度。
			// 注释的写法是获取左右扩展高度后，形成矩形宽度。
			// leftExpandWidth := 0
			// if len(indexStack) == 0 {
			//   leftExpandWidth = shortBoardIndex
			// } else {
			//	 leftExpandWidth = shortBoardIndex - indexStack[len(indexStack)-1] - 1
			// }
			// rightExpandWidth := i - shortBoardIndex - 1
			// matrixWidth := leftExpandWidth + rightExpandWidth + 1
			matrixWidth := 0
			if len(indexStack) == 0 {
				matrixWidth = i
			} else {
				matrixWidth = i - indexStack[len(indexStack)-1] - 1
			}

			// 3.1.4 计算矩形面积。
			result = max(result, matrixHeight*matrixWidth)
		}
		indexStack = append(indexStack, i)
	}

	// 4. 返回。
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
