

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
