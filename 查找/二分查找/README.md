
## 二分查找
1. 简单二分查找
    - 从非递减数组 `nums` 中找出任意一个索引`index`，使得 `nums[index] == target`，`index` 不存在时返回 `-1`。
        ``` go
        func binarySearch(nums []int, target int) int {
            l, r := 0, len(nums)-1
            for l <= r {
                mid := l + (r-l)/2
                if nums[mid] == target {
                    return mid
                } else {
                    if nums[mid] > target {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
            }
            return -1
        }
        ```
2. 进阶二分查找
    - 从非递减数组 `nums` 中找出一个索引 `index`，使得 `nums[index] > target`，要求 `index` 尽可能的小，不存在时返回 `len(nums)`。
        ``` go
        // 例子: firstGreater([]int{1,2,3,4,5}, 4) -> return 4
        // 「第一个大于」
        func firstGreater(nums []int, target int) int {
            l, r := 0, len(nums)-1
            for l <= r {
                mid := l + (r-l)/2
                if nums[mid] == target {
                    l = mid + 1
                } else {
                    if nums[mid] > target {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
            }
            return l
        }
        ```
    - 从非递减数组 `nums` 中找出一个索引 `index`，使得 `nums[index] >= target`，要求 `index` 尽可能的小，不存在时返回 `len(nums)`。
        ``` go
        // 例子: firstGreaterOrEqual([]int{1,2,3,4,5}, 4) -> return 3
        // 「第一个大于等于」
        func firstGreaterOrEqual(nums []int, target int) int {
            l, r := 0, len(nums)-1
            for l <= r {
                mid := l + (r-l)/2
                if nums[mid] == target {
                    r = mid - 1
                } else {
                    if nums[mid] > target {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
            }
            return l
        }
        ```
    - 从非递减数组 `nums` 中找出一个索引 `index`，使得 `nums[index] <= target`，要求 `index` 尽可能的大，不存在时返回 `-1`。
        ``` go
        // 例子: lastLessOrEqual([]int{1,2,3,4,5}, 4) -> return 3
        // 「最后一个小于等于」
        func lastLessOrEqual(nums []int, target int) int {
            l, r := 0, len(nums)-1
            for l <= r {
                mid := l + (r-l)/2
                if nums[mid] == target {
                    l = mid + 1
                } else {
                    if nums[mid] > target {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
            }
            return r
        }
        ```
    - 从非递减数组 `nums` 中找出一个索引 `index`，使得 `nums[index] < target`，要求 `index` 尽可能的大，不存在时返回 `-1`。
        ``` go
        // 例子: lastLess([]int{1,2,3,4,5}, 4) -> return 2
        // 「最后一个小于」
        func lastLess(nums []int, target int) int {
            l, r := 0, len(nums)-1
            for l <= r {
                mid := l + (r-l)/2
                if nums[mid] == target {
                    r = mid - 1
                } else {
                    if nums[mid] > target {
                        r = mid - 1
                    } else {
                        l = mid + 1
                    }
                }
            }
            return r
        }
        ```
3. 框架
    ``` go
    func function(nums []int, target int) int {
        l, r := 0, len(nums)-1
        for l <= r {
            mid := l + (r-l)/2
            if nums[mid] == target {
                /* 
                    以下是「索引选择」，有3个选项，单选。 
                    
                    A. 找等于目标值的索引时: return mid
                    B. 找尽可能小的索引时: r = mid - 1
                    C. 找尽可能大的索引时: l = mid + 1
                */
            } else {
                if nums[mid] > target {
                    r = mid - 1
                } else {
                    l = mid + 1
                }
            }
        }
        /* 
            以下是「返回值选择」，有3个选项，单选。

            A. 找等于目标值的索引时: return -1
            B. 找小于、小于等于target的索引时: return r
            C. 找大于、大于等于target的索引时: return l
        */
    }
    
    ```
    
4. 注意
    - [ ] 以上查找要求数组非递减，所以不适用于旋转排序数组。
    - [ ] 解决旋转排序数组问题，需要用到二分查找的思想。
    - [ ] 一般情况下，题目不会这么"裸"，所以实际应用中，可能需要修改一些地方、或进行一些转换。
5. 练习题
    - [ ] [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
    - [ ] [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)
    - [ ] [74. 搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)
    - [ ] [81. 搜索旋转排序数组 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/)
    - [ ] [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)
    - [ ] [154. 寻找旋转排序数组中的最小值 II](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)
    - [ ] [240. 搜索二维矩阵 II](https://leetcode-cn.com/problems/search-a-2d-matrix-ii/submissions/)
    - [ ] [278. 第一个错误的版本](https://leetcode-cn.com/problems/first-bad-version/)
    - [ ] [704. 二分查找](https://leetcode-cn.com/problems/binary-search/)

