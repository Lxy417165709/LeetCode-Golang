package 数学

// ------------------ O(n)时间复杂度,O(n)空间复杂度的代码 -----------------
func kthFactor(n int, k int) int {
	factors := getSortedFactors(n)
	if k-1 >= len(factors) {
		return -1
	}
	return factors[k-1]
}

func getSortedFactors(n int) []int {
	factors := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

// ------------------ O(sqrt(n))时间复杂度,O(n)空间复杂度的代码 -----------------
func kthFactor(n int, k int) int {
	factors := getSortedFactors(n)
	if k-1 >= len(factors) {
		return -1
	}
	return factors[k-1]
}

func getSortedFactors(n int) []int {
	factors := make([]int, getCountOfFactor(n))
	indexOfLeftInsert, indexOfRightInsert := 0, len(factors)-1
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors[indexOfLeftInsert] = i
			factors[indexOfRightInsert] = n / i
			indexOfLeftInsert++
			indexOfRightInsert--
		}
	}
	return factors
}

func getCountOfFactor(n int) int {
	countOfFactor := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				countOfFactor++
			} else {
				countOfFactor += 2
			}
		}
	}
	return countOfFactor
}

// ------------------ O(sqrt(n))时间复杂度, O(1)空间复杂度的代码 -----------------
func kthFactor(n int, k int) int {
	factor := -1
	if isFactorInLeftPart(n, k) {
		factor = getFactorInLeftPart(n, k)
	} else {
		factor = getFactorInRightPart(n, k)
	}
	return factor
}

func isFactorInLeftPart(n int, k int) bool {
	orderOfMidFactor := (getCountOfFactor(n) + 1) / 2
	return k <= orderOfMidFactor
}

func getFactorInRightPart(n int, rightK int) int {
	leftK := getCountOfFactor(n) + 1 - rightK
	leftFactor := getFactorInLeftPart(n, leftK)
	if leftFactor == -1 {
		return -1
	}
	return n / leftFactor
}

func getFactorInLeftPart(n int, leftK int) int {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			leftK--
			if leftK == 0 {
				return i
			}
		}
	}
	return -1
}

func getCountOfFactor(n int) int {
	countOfFactor := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i == n {
				countOfFactor++
			} else {
				countOfFactor += 2
			}
		}
	}
	return countOfFactor
}

/*
	题目链接: https://leetcode-cn.com/problems/the-kth-factor-of-n/
	总结
		1. 最开始，这题我采用了最笨的方法AC，时空复杂度都为O(n)
		2. 可以把求因子的时间复杂度降低为O(sqrt(n)),空间复杂度也可以优化为O(1), 详情看题解。
*/
