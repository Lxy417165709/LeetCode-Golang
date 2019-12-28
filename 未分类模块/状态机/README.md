# 简单状态机 (TODO)
## 1. 框架
```go

// 操作 -> 数字 映射函数
func hashOpt(opt uint8) int {
	/* 
	    1. 将操作映射为数字，使用数字表示操作。
	    (使用switch结构可以很直观的将操作映射为哈希)
	*/
}

// 调用状态机的函数
func DFACaller() {
	/* 2. 通过分析，构建状态机矩阵 */ 
	/* 3. 通过分析，构建终止状态向量 */ 
	/* 4. 确定当前状态 */                
	/* 5. 执行状态转换 */
	/* 6. 返回最终结果 */
}
```

## 2. 实例
### 2.1 检测大写字母
[传送门](https://leetcode-cn.com/problems/detect-capital/)
```go
// 操作 -> 数字 映射函数
func hashOpt(opt uint8) int {

	// 操作映射表  (用数字表示操作)
	// 小写字母: 0
	// 大写字母: 1
	// 非法输入: 2 (本题没有这操作)

	/* 1. 将操作映射为数字，使用数字表示操作。 */
	//  假如题目有 n 种操作，那么返回值应该有 n 种。
	//	操作与数字的映射关系可以自定义。
	//	比如: [小写字母操作] 可以映射为任何数字，只要不发生冲突，
	//	即: 假如 [小写字母操作] 映射为 (数字1)，那么 [大写字母操作] 就不能映射为 (数字1)。
	switch {
	case opt >= 'a' && opt <= 'z':
		return 0	// 第 1 种返回值
	case opt >= 'A' && opt <= 'Z':
		return 1	// 第 2 种返回值
	default:
		return 2	// 题目说输入都是合法的，所以这里并不会执行，所以不算一种返回值。
	}
}

// DFA调用者
func detectCapitalUse(s string) bool {

	// 状态映射表 (用数字表示状态)
	// 空白态:                 0
	// 小写态:                 1
	// 单个大写字母态:         2
	// 多个大写字母态:         3
	// 非法态:                 4

	/* 2. 通过分析，构建状态机矩阵 */
	// 假如有 m 种状态, n 种操作，那么状态机矩阵就是 m*n 的矩阵
	matrixOfDFA := [][]int{
		{1, 2}, // 空白态     		    「状态0」
		{1, 4}, // 小写态     		    「状态1」
		{1, 3}, // 单个大写字母态	    「状态2」
		{4, 3}, // 多个大写字母态	    「状态3」
		{4, 4}, // 非法态	            「状态4」

	}
	// 解释:
	// 		(1) 第 1 行 {1, 2} 表示:
	// 				空白态 -- 输入小写字母 --> 小写态。              
	// 				(「状态0」 -- [操作0] --> 「状态1」)
	//
	// 				空白态 -- 输入大写字母 --> 单个大写字母态。 
	// 				(「状态0」 -- [操作1] --> 「状态2」)
	//
	// 		(2) 第 4 行 {4, 3} 表示:
	// 				多个大写字母态 -- 输入小写字母 --> 非法态。   			
	// 				(「状态3」 -- [操作0] --> 「状态4」)
	
	// 				多个大写字母态 -- 输入大写字母 --> 多个大写字母态。   	
	// 				(「状态3」 -- [操作1] --> 「状态3」)

	/* 3. 通过分析，构建终止状态向量 */
	// 假如有 m 种状态, 那么终止向量的列数为 m。
	// 状态与数字的映射关系也可以自定义
	finishStates := []bool{true, true, true, true, false}
	// 			{true,     true,     true,     true, 	 false}
	//			「状态0」，「状态1」，「状态2」，「状态3」，「状态4」
	// 解释:
	//		(1) 题目允许空字符串(空白态) 作为结尾:             所以 finishStates[0] == true。 	(「0」 指「状态0」)
	//		(2) 题目允许全是小写的字符串(小写态) 作为结尾: 	所以 finishStates[1] == true。 	(「1」 指「状态1」)
	//		(3) 题目不允许 小写字母后加大写字母(非法态) 和
	//	            多个大写字母后加小写字母(非法态) 作为结尾: 	所以 finishStates[4] == false。	(「4」 指「状态4」)

	/* 4. 确定当前状态 */
	// 什么都没输入的时候，该字符串状态为: 空白态，对应的状态是「状态0」。
	nowState := 0

	/* 5. 执行状态转换 */
	for i := 0; i < len(s); i++ {
		nowState = matrixOfDFA[nowState][hashOpt(s[i])]
	}

	/* 6. 返回最终结果 */
	return finishStates[nowState]
}
```

### 2.2 有效数字
[传送门](https://leetcode-cn.com/problems/valid-number/submissions/)
```go
// 操作 -> 数字 映射函数
func hashOpt(opt uint8) int {
	/* 1. 将操作映射为数字，使用数字表示操作。 */
	switch {
	case opt >= '0' && opt <= '9':
		return 0
	case opt == 'e':
		return 1
	case opt == '+' || opt == '-':
		return 2
	case opt == '.':
		return 3
	case opt == ' ':
		return 4
	default:
		return 5
	}
}

// 调用状态机的函数
func isNumber(s string) bool {
	/* 2. 通过分析，构建状态机矩阵 */
	matrixOfDFA := [][]int{
		{1, 10, 8, 11, 0, 10},
		{1, 3, 10, 2, 6, 10},
		{7, 3, 10, 10, 6, 10},
		{5, 10, 4, 10, 10, 10},
		{5, 10, 10, 10, 10, 10},
		{5, 10, 10, 10, 6, 10},
		{10, 10, 10, 10, 6, 10},
		{7, 3, 10, 10, 6, 10},
		{9, 10, 10, 11, 10, 10},
		{9, 3, 10, 2, 6, 10},
		{10, 10, 10, 10, 10, 10},
		{7, 10, 10, 10, 10, 10},
	}
	
	/* 3. 通过分析，构建终止状态向量 */
	finishStates := []bool{false, true, true, false, false, true, true, true, false, true, false, false}

	/* 4. 确定当前状态 */
	nowState := 0

	/* 5. 执行状态转换 */
	for i := 0; i < len(s); i++ {
		nowState = matrixOfDFA[nowState][hashOpt(s[i])]
	}

	/* 6. 返回最终结果 */
	return finishStates[nowState]
}
```
## 3. 注意点
- [ ] DFA最复杂的地方在于构建状态机。我们可以先使用画图工具来画出状态转移图后，再得出状态机矩阵。
- [ ] 上面的状态命名可能有些不准确，我只是表达大致的意思 (语文不好真xxx)...
## 3. 练习题
- [ ] [65. 有效数字](https://leetcode-cn.com/problems/valid-number/submissions/)
- [ ] [520. 有效数字](https://leetcode-cn.com/problems/detect-capital/)