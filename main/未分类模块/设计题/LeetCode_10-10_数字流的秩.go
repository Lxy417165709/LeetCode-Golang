package main

type StreamRank struct {
	orderedArray []int
}

func Constructor() StreamRank {
	return StreamRank{make([]int, 0)}
}

func (sr *StreamRank) Track(x int) {
	sr.insertNumber(x)
	sr.insertSort()
}
func (sr *StreamRank) GetRankOfNumber(x int) int {
	return sr.getTheIndexOfFirstNumberGreaterThanReference(x)
}

func (sr *StreamRank) insertNumber(element int) {
	sr.orderedArray = append(sr.orderedArray, element)
}
func (sr *StreamRank) insertSort() {
	for i := len(sr.orderedArray) - 1; i >= 1; i-- {
		if sr.orderedArray[i-1] > sr.orderedArray[i] {
			sr.orderedArray[i-1], sr.orderedArray[i] = sr.orderedArray[i], sr.orderedArray[i-1]
		} else {
			break
		}
	}
}

// 涉及算法：二分查找
func (sr *StreamRank) getTheIndexOfFirstNumberGreaterThanReference(reference int) int {
	leftBoard, rightBoard := 0, len(sr.orderedArray)-1
	for leftBoard <= rightBoard {
		readingIndex := (leftBoard + rightBoard) / 2
		readingNumber := sr.orderedArray[readingIndex]
		if readingNumber == reference {
			leftBoard = readingIndex + 1
		} else {
			if readingNumber > reference {
				rightBoard = readingIndex - 1
			} else {
				leftBoard = readingIndex + 1
			}
		}
	}
	return leftBoard
}

