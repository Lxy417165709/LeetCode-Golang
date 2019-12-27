
## 求最大子串和、最长递增子串长度
### 1. 框架
```go
func DP() {
    /* 1. 明白dp[i]定义后，初始化dp数组 */
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            /* 2. dp[i]基础情况处理 (指: nums[i]前面没有元素时) */
            continue
        }
        /* 3. 根据条件更新dp[i] */
    }
}
```


### 2. 实例
#### 2.1 求最大子串和 (原问题描述: 求最大连续子序列和)
[传送门](https://leetcode-cn.com/problems/maximum-subarray/)
```go
const inf = 100000000000 // 要定义的足够大，一般要大于int32

func maxSubArray(nums []int) int {
	/* 1. 明白dp[i]定义后，初始化dp数组 */
	dp := make([]int, len(nums)) // 定义dp[i]为: 以nums[i]为结尾的最大子串和

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			/* 2. dp[i]基础情况处理 (指: nums[i]前面没有元素时) */
			dp[i] = nums[i]
			continue
		}
		/* 3. 根据条件更新dp[i] */
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		// 「步骤3」可以简化为:  dp[i] = max(nums[i], dp[i-1] + nums[i])
		
	}
	ans := -inf
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
#### 2.2 求最长递增子串长度 (原问题描述: 求最长连续递增子序列长度)
[传送门](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/submissions/)
```go
func findLengthOfLCIS(nums []int) int {
	/* 1. 明白dp[i]定义后，初始化dp数组 */
	dp := make([]int, len(nums)) // 定义dp[i]为: 以nums[i]为结尾的最长递增串长度
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			/* 2. dp[i]基础情况处理 (nums[i]前面没有元素时) */
			dp[i] = 1
			continue
		}
		/* 3. 根据条件更新dp[i] */
		if nums[i] > nums[i-1]{
			dp[i] = dp[i-1] + 1
		}else{
			dp[i] = 1
		}
	}
	maxLength := 0
	for i:=0;i<len(dp);i++{
		maxLength = max(maxLength,dp[i])
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
### 3. 练习题
- [ ] [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)
- [ ] [674. 最长连续递增序列](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/submissions/)
