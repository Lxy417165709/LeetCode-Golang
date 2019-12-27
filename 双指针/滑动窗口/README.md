# 滑动窗口 (TODO)
## 1. 框架
```go
func movingWindow() {
    /* 1. 初始化窗口数据结构，用于记录窗口内的信息 */
    first, last := 0, 0    // 窗口的左右边界
    for last < len(s) {
        /* 2. 把 last 指向的元素加入窗口 */
        for first < len(s) && /* 3. 判断当前窗口内的元素是否符合条件 */ {
            // 不满足时
            /* 4. 把 first 指向的元素移出窗口 */
            /* 5.a 在这写更新窗口最大值的代码 */
            first++
        }
        /* 5.b 在这写更新窗口最大值的代码 */
        last++
    }
}
```
## 2. 实例
### 2.1 无重复字符的最长子串
[传送门](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
```go
func lengthOfLongestSubstring(s string) int {
    
    /* 1. 初始化窗口数据结构，用于记录窗口内的信息 */
    count := make(map[uint8]int) // 用来记录窗口中的每个字符出现了多少次
    first, last := 0, 0          // 窗口的左右边界
    
    maxLength := 0
    for last < len(s) {
        /* 2. 把 last 指向的元素加入窗口 */
        count[s[last]]++
    
        /* 3. 判断当前窗口内的元素是否符合条件 */
        for first < len(s) && count[s[last]] >= 2 {
        	// 不满足时
            /* 4. 把 first 指向的元素移出窗口 */
            count[s[first]]--
            first++
        }
        /* 5.b 在这写更新窗口最大值的代码 */
        maxLength = max(maxLength, last-first+1)
        last++
    }
    return maxLength
}
```
   
## 3. 练习题
- [ ] [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
- [ ] [209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)