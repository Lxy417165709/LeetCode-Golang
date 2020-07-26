package 未归类

func hanota(A []int, B []int, C []int) []int {
	move(len(A), &A, &B, &C)
	return C
}

// 将 A 上面的 n 个盘子，借助 B，移动到 C 中 (移动的每步要符合汉诺塔规则)
func move(n int, A *[]int, B *[]int, C *[]int) {
	if n == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		return
	}
	move(n-1, A, C, B)
	move(1, A, B, C)
	move(n-1, B, A, C)
}

/*
	题目链接: https://leetcode-cn.com/problems/hanota-lcci/
	总结
		1. 这题我还不太理解，不太理解的点在参数n --- n是否可以省去呢？
		2. 这题不传引用的话，会出现错误。 (不传引用会导致 append 后，函数内的参数修改了，调用者的传入的参数却没改)
*/
