
## 求最长上升、整除、定差子序列的长度
### 1. 框架
```go
func DP() {
    /* 1. 明白dp[i]定义后，初始化dp数组 */
    for i := 0; i < len(nums); i++ {
    	/* 2. dp[i]基础情况处理 (指: 子序列只有nums[i]时) */
        for t := 0; t < i; t++ {
            if /* 3. 判断nums[i]与nums[t]的关系，决定是否更新dp[i] */ {
                /* 4. 更新dp[i] */
            }
        }
    }
}
```
「步骤3」的常用判断条件举例:
- `nums[i] > nums[t]` 用以更新 最长上升子序列。
- `nums[i] % nums[t] == 0` 用以更新 最长整除子序列。 
- `nums[i] - nums[t] == difference` 用以更新 最长定差子序列。 (`difference` 指最长定差子序列的公差)



### 2. 实例
#### 2.1 求最长上升子序列的长度
[传送门](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
```go
func lengthOfLIS(nums []int) int {
    /* 1. 明白dp[i]定义后，初始化dp数组 */
    dp := make([]int, len(nums))    // 这里定义dp[i]为: 以nums[i]为结尾的最长上升子序列长度
   
    for i := 0; i < len(nums); i++ {
        /* 2. dp[i]基础情况处理 (指: 子序列只有nums[i]时) */
        dp[i] = 1
        for t := 0; t < i; t++ {
            /* 3. 判断nums[i]与nums[t]的关系，决定是否更新dp[i] */
            if nums[i] > nums[t] {
                /* 4. 更新dp[i] */
                dp[i] = max(dp[i], dp[t]+1)
            }
        }
    }
    ans := 0
    for i := 0; i < len(dp); i++ {
        ans = max(dp[i], ans)
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
#### 2.2 求最大整除子集
[传送门](https://leetcode-cn.com/problems/largest-divisible-subset/submissions/)
```go
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	/* 1. 明白dp[i]定义后，初始化dp数组 */
	dp := make([][]int, len(nums))   // 这里定义dp[i]为: 以nums[i]为结尾的最大整除子集

	for i := 0; i < len(nums); i++ {
		/* 2. dp[i]基础情况处理 (指: 子序列只有nums[i]时) */
		dp[i] = append(dp[i], nums[i])
		for t := 0; t < i; t++ {
			/* 3. 判断nums[i]与nums[t]的关系，决定是否更新dp[i] */
			if nums[i]%nums[t] == 0 {
				/* 4. 更新dp[i] */
				dp[i] = max(dp[i], append(newSlice(dp[t]), nums[i]))
			}
		}
	}

	ans := make([]int, 0)
	for i := 0; i < len(dp); i++ {
		ans = max(dp[i], ans)
	}
	return ans
}

// 返回长度最长的切片
func max(a, b []int) []int {
	if len(a) > len(b) {
		return a
	}
	return b
}

// 切片深拷贝
func newSlice(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}
```

### 3. 拓展问题
- [ ] 如果要求的是最长递减、最长非递减子序列的长度，那应该怎么做呢？

### 4. 练习题
- [ ] [300. 最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
- [ ] [368. 最大整除子集](https://leetcode-cn.com/problems/largest-divisible-subset/submissions/)
- [ ] [673. 最长递增子序列的个数](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/)
- [ ] [1218. 最长定差子序列](https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/)  (这里只是为这题提供了一种dp思路，实际上还有更快速的方法)
