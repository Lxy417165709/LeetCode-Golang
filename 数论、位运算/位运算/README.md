
## 位运算
1. 位运算性质
    - 异或运算
        - 满足交换律: `x ^ y == y ^ x`
        - 满足结合律: `(x ^ y) ^ z == x ^ (y ^ z)`
        - 相同的数异或为 `0`: `x ^ x == 0`
        - 任何数和 `0` 异或都为本身: `x ^ 0 == x`
    - ...
    
2. 小技巧
    - 获取 `整数x` 的最低位`1`的对应的值: `x & (-x)`。
    - 去除 `整数x` 的最低位`1`后的值: `x & (x-1)`。
    - 实现 `整数x` 与 `整数y` 的不进位加法: `x ^ y`
    - 实现 `0` 与 `1`的相互转换: 
        - `0 ^ 1 == 1`
        - `1 ^ 1 == 0`
    - 关于取模
        - 获取 `整数x` 模`2`后的值: `x & 1`
        - 获取 `整数x` 模`4`后的值: `x & 3`
        - 获取 `整数x` 模`8`后的值: `x & 7`
        - ...

3. 拓展
    - 判断`整数x`是否为`2`的幂
        ```go
        func isPowerOfTwo(n int) bool {
            return n > 0 && n&(n-1) == 0
        }
        ```
    - 判断`整数x`是否为`4`的幂
        ```go
        // 只适用于32位整数
        func isPowerOfFour(num int) bool {
            return isPowerOfTwo(num) && num&0x55555555 != 0
        }
        
        // 适用于32、64位整数，但不适用于32位系统
        func isPowerOfFour(num int) bool {
            return isPowerOfTwo(num) && num&0x5555555555555555 != 0
        }
        
        // 适用于32、64位整数，适用于32、64位系统
        func isPowerOfFour(num int) bool {
            return isPowerOfTwo(num) && num%3 != 0
        }
        
        func isPowerOfTwo(n int) bool {
            return n > 0 && n&(n-1) == 0
        }
        ```
    - 利用位运算实现`整数x`、`整数y`的交换。
        ```go
        // 这里只是形参的交换，如果要实现实参的交换，那么可以传入指针。
        func swap(x,y int) {
            x = x ^ y
            y = x ^ y
            x = x ^ y
        }
        ```
    - 利用位运算获取`整数x`有多少个`1`。
        ```go
        func hammingWeight(num int) int {
            count := 0
            for num != 0 {
                count++
                num = num & (num - 1) 
            }
            return count
        }
        ```
    - 计算 `整数x`、`整数y`的有多少个不同的二进制位。
        ```go
        func hammingDistance(x,y int) int {
            return hammingWeight(x^y)
        }
        
        func hammingWeight(num int) int {
            count := 0
            for num != 0 {
                count++
                num = num & (num - 1)
            }
            return count
        }
        ```
    - 利用位运算实现不用 `+` 做加法。
        ``` go
        func getSum(a int, b int) int {
            sum, cnumsy := a^b, (a&b)<<1
            if cnumsy == 0 {
                return sum
            }
            return getSum(sum, cnumsy)
        }
        ```
    - 位运算还可以实现小容量的`map`。

4. 注意
    - 以上位运算基于补码。

5. 练习题
    - [ ] [136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/)
    - [ ] [191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)
    - [ ] [231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)
    - [ ] [342. 4的幂](https://leetcode-cn.com/problems/power-of-four/)
    - [ ] [371. 两整数之和](https://leetcode-cn.com/problems/sum-of-two-integers/submissions/)
    - [ ] [461. 汉明距离](https://leetcode-cn.com/problems/hamming-distance/)
