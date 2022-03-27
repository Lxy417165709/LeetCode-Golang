package 一维数组

import "fmt"

// --------------------------------------------------- 1. 动态规划(开始) ---------------------------------------------------

// trap 接雨水。
// 动态规划解法: 计算所有水柱高度的总和。 (垂直)
// 水柱高度: 向左最大高度、向右最大高度的最小值 - 当前位置的高度。
func trap(heights []int) int {
	// 1. 构建当前点向左的最大值映射。
	maxHeightToLeft := make([]int, len(heights)+2) // maxHeightToLeft[pos] 表示 heights[:pos] 的最大值。
	for i := 0; i < len(heights); i++ {
		pos := i + 1
		maxHeightToLeft[pos] = max(maxHeightToLeft[pos-1], heights[i])
	}

	// 2. 构建当前点向右的最大值映射。
	maxHeightToRight := make([]int, len(heights)+2) // maxHeightToRight[pos] 表示 heights[pos-1:] 的最大值。
	for i := len(heights) - 1; i >= 0; i-- {
		pos := i + 1
		maxHeightToRight[pos] = max(maxHeightToRight[pos+1], heights[i])
	}

	// 3. 获取雨水量。
	rain := 0
	for i := 0; i < len(heights); i++ {
		pos := i + 1
		rain += min(maxHeightToLeft[pos], maxHeightToRight[pos]) - heights[i]
	}

	// 4. 返回。
	return rain
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// --------------------------------------------------- 1. 动态规划(结束) ---------------------------------------------------

// --------------------------------------------------- 2. 最小栈(开始) ---------------------------------------------------

// trap 接雨水。
// 最小栈解法: 计算所有水滩容量的总和。 (水平)
// 水滩高度: (左侧第一个比滩底高的位置 与 右侧第一个比滩底高的位置 的距离) * (左右侧短板高度 - 滩底高度)
// PS: 使用单调递减栈、单调非递增栈都能解决该题。
//     切换方式是 heights[indexStack[len(indexStack)-1]] <= heights[i] 改为
//              heights[indexStack[len(indexStack)-1]] < heights[i]。
func trap(heights []int) int {
	// 1. 初始化。
	indexStack := make([]int, 0) // 单调递减栈。(大小关系指的是值，该栈内部存放指向值的索引)
	rain := 0

	// 2. 处理。
	for i := 0; i < len(heights); i++ {
		// 2.1 累加雨水量、出栈。
		for len(indexStack) != 0 && heights[indexStack[len(indexStack)-1]] <= heights[i] {
			// 2.2.1 大于 1，说明存在右边界。
			if len(indexStack) > 1 {
				indexOfLeftBound, indexOfRightBound := indexStack[len(indexStack)-2], i                         // 水滩的左右边界。
				indexOfStackTop := indexStack[len(indexStack)-1]                                                // 滩底高度。
				height := min(heights[indexOfLeftBound], heights[indexOfRightBound]) - heights[indexOfStackTop] // 水滩高度 = 左右侧短板高度 - 滩底高度。
				width := i - indexOfLeftBound - 1                                                               // 水滩宽度 = 左侧第一个比滩底高的位置 与 右侧第一个比滩底高的位置 的距离。
				rain += height * width
			}

			// 2.2.1 出栈。
			indexStack = indexStack[:len(indexStack)-1]
		}

		// 2.2 入栈。
		indexStack = append(indexStack, i)
	}

	// 3. 返回。
	return rain
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// --------------------------------------------------- 2. 最小栈(结束) ---------------------------------------------------
