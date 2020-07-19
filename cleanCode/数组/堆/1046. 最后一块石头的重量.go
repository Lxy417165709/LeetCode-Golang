package 堆

import "sort"

// ---------------------------- 暴力解法 ----------------------------
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度 O(n^2 * log_n)
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		if stones[len(stones)-1] == stones[len(stones)-2] {
			stones = stones[:len(stones)-2]
		} else {
			stones[len(stones)-2] = stones[len(stones)-1] - stones[len(stones)-2]
			stones = stones[:len(stones)-1]
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

// ---------------------------- 采用大顶堆 ----------------------------
// 执行用时：0 ms, 在所有 Go 提交中击败了 100.00% 的用户
// 内存消耗：2 MB, 在所有 Go 提交中击败了 100.00% 的用户
//
// 时间复杂度 O(n * log_n)
func lastStoneWeight(stones []int) int {
	heap := NewHeapBuildingOnNums(stones)
	for heap.GetSize() >= 2 {
		firstBigWeight := heap.Pop()
		secondBigWeight := heap.Pop()
		if firstBigWeight > secondBigWeight {
			heap.Push(firstBigWeight - secondBigWeight)
		}
	}
	if heap.GetSize() == 0 {
		return 0
	}
	return heap.Pop()
}

// ------------------- Heap -------------------
type Heap struct {
	data           []int
	indexOfLastNum int
}

func NewHeapBuildingOnNums(nums []int) *Heap {
	heap := NewHeap(len(nums))
	for i := 0; i < len(nums); i++ {
		heap.Push(nums[i])
	}
	return heap
}

func NewHeap(maxSize int) *Heap {
	return &Heap{
		data:           make([]int, maxSize*2+2),
		indexOfLastNum: 0,
	}
}

func (h *Heap) GetTop() int {
	return h.data[1]
}

func (h *Heap) GetSize() int {
	return h.indexOfLastNum
}

func (h *Heap) IsEmpty() bool {
	return h.GetSize() == 0
}

func (h *Heap) Pop() int {
	h.indexOfLastNum--
	top := h.data[1]
	h.swap(1, h.indexOfLastNum+1)
	h.downAdjust(1, h.indexOfLastNum)
	return top
}

func (h *Heap) Push(val int) {
	h.indexOfLastNum++
	h.data[h.indexOfLastNum] = val
	h.upAdjust(h.indexOfLastNum)
}

func (h *Heap) upAdjust(indexOfAdjustNum int) {
	curIndex := indexOfAdjustNum
	for {
		if h.isFatherNotExist(curIndex) {
			return
		}
		fatherIndex := h.getIndexOfFather(curIndex)
		if h.data[fatherIndex] >= h.data[curIndex] {
			return
		}
		h.swap(fatherIndex, curIndex)
		curIndex = fatherIndex
	}
}

func (h *Heap) downAdjust(lowIndex, highIndex int) {
	curIndex := lowIndex
	for {
		if h.isSonNotExist(curIndex, highIndex) {
			return
		}
		indexOfMaxValueSon := h.getIndexOfMaxValueSon(curIndex, highIndex)
		if h.data[indexOfMaxValueSon] <= h.data[curIndex] {
			return
		}
		h.swap(curIndex, indexOfMaxValueSon)
		curIndex = indexOfMaxValueSon
	}
}

func (h *Heap) isFatherNotExist(curIndex int) bool {
	return curIndex/2 == 0
}

func (h *Heap) getIndexOfFather(curIndex int) int {
	return curIndex / 2
}

func (h *Heap) isSonNotExist(curIndex int, highIndex int) bool {
	leftSonIndex := curIndex * 2
	return leftSonIndex > highIndex
}

func (h *Heap) getIndexOfMaxValueSon(curIndex int, highIndex int) int {
	indexOfMaxValueSon := curIndex
	leftSonIndex, rightSonIndex := curIndex*2, curIndex*2+1
	if leftSonIndex <= highIndex {
		indexOfMaxValueSon = leftSonIndex
	}
	if rightSonIndex <= highIndex && h.data[rightSonIndex] > h.data[leftSonIndex] {
		indexOfMaxValueSon = rightSonIndex
	}
	return indexOfMaxValueSon
}

func (h *Heap) swap(index1, index2 int) {
	h.data[index1], h.data[index2] = h.data[index2], h.data[index1]
}

/*
	题目链接: https://leetcode-cn.com/problems/last-stone-weight/
*/
