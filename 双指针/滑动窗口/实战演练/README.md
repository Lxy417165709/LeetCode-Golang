## 滑动窗口 (TODO)
1. 框架
    ```go
    func movingWindow() {
        /* 1. 初始化窗口数据结构，用于记录窗口内的信息 */
        first, last := 0, 0    // 窗口的左右边界
        for last < len(s) {
            /* 2. 把 last 指向的元素加入窗口 */
            for first < len(s) && /* 3. 判断条件 */ {
                /* 4. 把 first 指向的元素移出窗口 */
                /* 5.a 如果要求的是窗口最小值，那么在这里更新窗口最小值 */
                first++
            }
            /* 5.b 如果要求的是窗口最大值，那么在这里更新窗口最大值 */
            last++
        }
    }
    ```
2. 实例 ( [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/) )

    ```go
    /*
       给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
    */
    func lengthOfLongestSubstring(s string) int {
    	maxLength := 0
 	    /* 1. 初始化窗口数据结构，用于记录窗口内的信息 */
    	count := make(map[uint8]int) // 用来记录窗口中的每个字符出现了多少次
 	
    	first, last := 0, 0          // 窗口的左右边界
    	for last < len(s) {
 		    /* 2. 把 last 指向的元素加入窗口 */
    		count[s[last]]++
 		
 		    /* 3. 判断条件 */
    		for first < len(s) && count[s[last]] >= 2 {
 			    /* 4. 把 first 指向的元素移出窗口 */
    			count[s[first]]--

    			first++
    		}
 		    /* 5.b 如果要求的是窗口最大值，那么在这里更新窗口最大值 */
    		maxLength = max(maxLength, last-first+1)
    		last++
    	}
    	return maxLength
    }
    ```
   
3. 练习题
    - [ ] [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
    - [ ] [209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)