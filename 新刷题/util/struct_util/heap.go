package struct_util

// ------------------------------------------ 1. 堆(开始) ------------------------------------------

// MyHeap 堆。
type MyHeap struct {
	data         []interface{}                       // 堆数据。
	compareValue func(interface{}, interface{}) bool // 堆内保证: compareValue(堆顶元素, 非堆顶元素) 必为 true。
	lastIndex    int                                 // 堆内最后的索引。
}

// NewMyHeap 建堆。
func NewMyHeap(cap int, compare func(interface{}, interface{}) bool) *MyHeap {
	return &MyHeap{
		data:         make([]interface{}, cap*2+2),
		compareValue: compare,
		lastIndex:    0,
	}
}

// Push 将元素推入堆。
func (m *MyHeap) Push(obj interface{}) {
	// 1. 维护索引。
	m.lastIndex++

	// 2. 写入元素。
	m.data[m.lastIndex] = obj

	// 3. 向上调整元素位置。
	m.upAdjust(m.lastIndex)
}

// Pop 推出堆顶元素。
func (m *MyHeap) Pop() interface{} {
	// 1. 异常返回。
	if m.Size() == 0 {
		panic("heap is empty.")
	}

	// 2. 存储原堆顶元素。
	top := m.data[1]

	// 3. 维护索引。
	m.lastIndex--

	// 4. 将尾元素移到堆顶，向下调整该元素的位置。
	m.swap(1, m.lastIndex+1)
	m.downAdjust(1)

	// 5. 返回原堆顶元素。
	return top
}

// Top 获取堆顶元素。
func (m *MyHeap) Top() interface{} {
	// 1. 异常返回。
	if m.Size() == 0 {
		panic("heap is empty.")
	}

	// 2. 返回堆顶元素。
	return m.data[1]
}

// Size 获取堆内元素个数。
func (m *MyHeap) Size() int {
	return m.lastIndex
}

// Sort 返回排序结果。 (会清空堆)
func (m *MyHeap) Sort() []interface{} {
	var result []interface{}
	for m.Size() > 0 {
		result = append(result, m.Pop())
	}
	return result
}

// upAdjust 向上调整。
func (m *MyHeap) upAdjust(index int) {
	for {
		// 1. 获取父节点索引。
		fatherIndex := index / 2

		// 2. 已是根节点，退出。
		if fatherIndex == 0 {
			break
		}

		// 3. 父子节点已满足堆条件，退出。
		if m.compare(fatherIndex, index) {
			break
		}

		// 4. 父子节点不满足堆条件，此时交换父子节点值，继续向上处理。
		m.swap(fatherIndex, index)
		index /= 2
	}
}

// upAdjust 向上调整。
func (m *MyHeap) downAdjust(index int) {
	for {
		// 1. 获取左右子节点索引。
		leftSonIndex, rightSonIndex := index*2, index*2+1

		// 2. 获取左右子节点中，权重最大的节点。
		indexOfMaxWeightNode := -1
		if leftSonIndex <= m.lastIndex && rightSonIndex <= m.lastIndex {
			// 2.1 存在左右子节点。
			if m.compare(leftSonIndex, rightSonIndex) {
				// 2.1.1 表示左节点权重大。
				indexOfMaxWeightNode = leftSonIndex
			} else {
				// 2.1.2 表示右节点权重大。
				indexOfMaxWeightNode = rightSonIndex
			}
		} else if leftSonIndex <= m.lastIndex {
			// 2.1 只存在左节点。
			indexOfMaxWeightNode = leftSonIndex
		}

		// 3. 不存在左右子节点，说明自身是叶子节点，退出。
		if indexOfMaxWeightNode == -1 {
			break
		}

		// 4. 如果左右子树权重最大的节点，不大于本节点权重，退出。
		if m.compare(index, indexOfMaxWeightNode) {
			break
		}

		// 5. 交换父子节点值，继续向下处理。
		m.swap(indexOfMaxWeightNode, index)
		index = indexOfMaxWeightNode
	}
}

// compare 比对。
func (m *MyHeap) compare(index1, index2 int) bool {
	return m.compareValue(m.data[index1], m.data[index2])
}

// swap 交换。
func (m *MyHeap) swap(index1, index2 int) {
	m.data[index1], m.data[index2] = m.data[index2], m.data[index1]
}

// ------------------------------------------ 1. 堆(结束) ------------------------------------------
