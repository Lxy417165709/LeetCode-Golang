package 获取

// ------------------------ 方法1: getSum 和 getTilt 函数分开 ------------------------
func findTilt(root *TreeNode) int {
	return getTilt(root)
}

func getTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return abs(getSum(root.Right)-getSum(root.Left)) + getTilt(root.Left) + getTilt(root.Right)
}

func getSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + getSum(root.Left) + getSum(root.Right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// ------------------------ 方法2: getSum 和 getTilt 函数合并 ------------------------
func findTilt(root *TreeNode) int {
	_, tilt := getSumAndTilt(root)
	return tilt
}

func getSumAndTilt(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftTreeSum, leftTreeTilt := getSumAndTilt(root.Left)
	rightTreeSum, rightTreeTilt := getSumAndTilt(root.Right)
	return root.Val + leftTreeSum + rightTreeSum, abs(leftTreeSum-rightTreeSum) + leftTreeTilt + rightTreeTilt
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

/*
	题目链接: https://leetcode-cn.com/problems/binary-tree-tilt/
*/
