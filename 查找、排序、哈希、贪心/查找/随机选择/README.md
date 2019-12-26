
## 随机选择
1. 选择数组中的第K小与第K大。
    ```go
    // 选择第k小的数 (重复的数次序不等)
    // 比如 [1 2 2 3]，2是第2小、也是第3小。
    func selectSmallKth(nums []int, k int) int {
    	l, r := 0, len(nums)-1
    	for l <= r {
    		index := randomPartition(nums, l, r)
    		if index+1 == k {
    			return nums[index]
    		} else {
    			if index+1 > k {
    				r = index - 1
    			} else {
    				l = index + 1
    			}
    		}
    	}
    	return -100000000 // 表示没找到 (k非法了)
    }
    
    // 选择第k大的数 (重复的数次序不等)
    func selectBigKth(nums []int, k int) int {
    	// 第k大 == 第 len(nums)-k+1 小
    	return selectSmallKth(nums, len(nums)-k+1)
    }


    // 随机划分 (l，r在合法范围内)
    // 返回的是索引值
    func randomPartition(nums []int, l int, r int) int {
        randomIndex := rand.Intn(r-l+1) + l                     // 选择随机索引
        nums[randomIndex], nums[l] = nums[l], nums[randomIndex] // 打乱数组
        guardIndex := l
        for l <= r {
            for l <= r && nums[l] <= nums[guardIndex] {
                l++
            }
            for l <= r && nums[r] >= nums[guardIndex] {
                r--
            }
            if l <= r {
                nums[l], nums[r] = nums[r], nums[l]
            }
        }
        nums[guardIndex], nums[r] = nums[r], nums[guardIndex]
        return r
    }
    ```
2. 拓展
    - [ ] 如果数组元素太多，以致内存无法装下，请问还有什么办法获得第K大吗？
3. 练习题
    - [ ] [414. 第三大的数](https://leetcode-cn.com/problems/third-maximum-number/)   
 