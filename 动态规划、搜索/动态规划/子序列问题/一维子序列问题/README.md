
## 求最长上升、整除、定差子序列的长度
1. 实例 ( 求最长上升子序列的长度 )
	```go
    func lengthOfLIS(nums []int) int {
   	    /* 1. 初始化操作 */
        dp := make([]int, len(nums))
       
        for i := 0; i < len(nums); i++ {
       	    /* 2. 赋值操作 */
            dp[i] = 1
            for t := 0; t < i; t++ {
           	    /* 3. 判断条件 */
                if nums[i] > nums[t] {
               	    /* 4. 执行具体操作 */
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
	```

2. 框架
	```go
     /* 1. 初始化操作 */
     for i := 0; i < len(nums); i++ {
         /* 2. 赋值操作 */
         for t := 0; t < i; t++ {
             if /* 3. 判断条件 */ {
                 /* 4. 执行具体操作 */
             }
         }
     }
	```
	- 「1. 初始化操作」 指: 初始化dp数组、及其它的数据结构(如果有)。
		- 例子: `dp := make([]int, len(nums))` 表示: 接下来dp数组存放的是整型数据且长度为`len(nums)`。
	- 「2. 赋值操作」 指: 当子序列仅仅包含`nums[i]`时，此时dp数组、其它的数据结构(如果有)应为什么值。
		- 例子: `dp[i] = 1` 表示: 当子序列仅仅包含`nums[i]`时，`dp[i] = 1`。
	- 「3. 判断条件」 指: 根据`nums[i]` 与 `nums[t]` 的关系，判断是否能拓展最长子序列。
		- 例子: `nums[i] > nums[t]` 拓展的是最长上升子序列。
		- 例子: `nums[i] % nums[t] == 0` 拓展的是最长整除子序列。 
		- 例子: `nums[i] - nums[t] == difference` 拓展的是最长定差子序列。 (`difference`指最长定差子序列的公差)
	- 「4. 具体操作」 指: 获得了以`nums[i]`为结尾的最长子序列后，接下来要做什么。

3. 拓展问题
    - [ ] 如果要求的是最长递减、最长非递减子序列的长度，那应该怎么做呢？
4. 练习题
    - [ ] [300. 最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)
    - [ ] [368. 最大整除子集](https://leetcode-cn.com/problems/largest-divisible-subset/submissions/)
    - [ ] [673. 最长递增子序列的个数](https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/)
    - [ ] [1218. 最长定差子序列](https://leetcode-cn.com/problems/longest-arithmetic-subsequence-of-given-difference/)  (这里只是为这题提供了一种dp思路，实际上还有更快速的方法)
