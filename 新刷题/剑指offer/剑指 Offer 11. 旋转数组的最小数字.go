package 剑指offer

func minArray(numbers []int) int {
	return getMinNum2(numbers, 0, len(numbers)-1)
}

// getMinNum1 获取最小值。 (第一版)
func getMinNum1(numbers []int, left, right int) int {
	// 1. 空返回。
	if left > right {
		return 1000000000
	}

	// 2. 只有一个元素，直接返回。
	if left == right {
		return numbers[left]
	}

	// 3. 初始化指针。
	mid := (left + right) / 2
	leftNum, rightNum := numbers[left], numbers[right]

	// 4. 左右端点相同，说明最小值在中间。
	if leftNum == rightNum {
		a := getMinNum1(numbers, left, mid)
		b := getMinNum1(numbers, mid+1, right)
		return min(a, b)
	}

	// 5. 左端点小于右端点，说明升序。
	if leftNum < rightNum {
		return leftNum
	}

	// 6. 右端点小于左端点，说明最小值在中间。
	a := getMinNum1(numbers, left, mid)
	b := getMinNum1(numbers, mid+1, right)
	return min(a, b)
}

// getMinNum2 获取最小值。 (第二版)
func getMinNum2(numbers []int, left, right int) int {
	// 1. 空返回。
	if left > right {
		return 1000000000
	}
	if left == right {
		return numbers[left]
	}

	// 2. 初始化指针。
	mid := (left + right) / 2
	rightNum := numbers[right]
	midNum := numbers[mid]

	// 3. 中值小于右值，最小值在左边。 (因为数组是升序的)
	if midNum < rightNum {
		return getMinNum2(numbers, left, mid)
	}

	// 4. 中值大于右值，最小值在右边。
	if midNum > rightNum {
		return getMinNum2(numbers, mid+1, right)
	}

	// 5. 中值等于右值，最小值可能在左右边，但可以排除右端点元素。
	return getMinNum2(numbers, left, right-1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// todo: 这里可以列下可能性。
// todo: 问题: 为什么只需要比对中值和右值，而不是左值和中值呢？
