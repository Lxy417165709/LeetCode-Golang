package 数组动态规划



// 状态: 位置、手中是否有数
// 操作: 取、不取
// 定义: 某状态下的最长湍流子数组长度

func maxTurbulenceSize(A []int) int {
	n,m := len(A),2
	if n==0{
		return 0
	}
	dp := newMatrix(n,m)
	for i:=0;i<len(A);i++{
		// 1. 处理 dp[i][0]
		if i >= 1{
			dp[i][0] = max(dp[i-1][0],dp[i-1][1])
		}

		// 2. 处理 dp[i][1]
		dp[i][1] = 1
		if i==1 && (A[i]>A[i-1] || A[i]<A[i-1]){
			dp[i][1] = max(dp[i][1],dp[i-1][1]+1)
		}
		if i>1{
			if A[i]>A[i-1]{
				if A[i-1]<A[i-2]{
					dp[i][1] = max(dp[i][1],dp[i-1][1]+1)
				}
				if A[i-1]>A[i-2]{
					dp[i][1] = max(dp[i][1],2)
				}
			}
			if A[i]<A[i-1]{
				if A[i-1]>A[i-2]{
					dp[i][1] = max(dp[i][1],dp[i-1][1]+1)
				}
				if A[i-1]<A[i-2]{
					dp[i][1] = max(dp[i][1],2)
				}
			}
		}
	}
	return max(dp[n-1][0],dp[n-1][1])

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newMatrix(n,m int) [][]int{
	matrix := make([][]int,n)
	for i:=0;i<len(matrix);i++{
		matrix[i] = make([]int,m)
	}
	return matrix
}
