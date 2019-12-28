
@[toc]

## 二叉树小总结
### 导言
1. 以下代码都存放于 [我的GitHub仓库](https://github.com/Lxy417165709/LeetCode-Golang) ，如果小伙伴觉得有用，请给我颗星星哈。
2. 以下代码都是提交过的，正确性可以保证。
### 代码
#### 属性获取
#####  1. 获取普通二叉树节点个数
   ```go
   func countNodes(root *TreeNode) int {
       if root == nil{
           return 0
       }
       return countNodes(root.Left) + countNodes(root.Right) + 1
   }
   ```
#####  2. 获取完全二叉树节点个数
```go
// 求完全二叉树节点个数
func countNodesExec(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := getCompleteBinaryTreeDepth(root.Left)
    rightDepth := getCompleteBinaryTreeDepth(root.Right)
    leftTreeNodes, rightTreeNodes := 0, 0

    if leftDepth == rightDepth {
        leftTreeNodes = getFullBinaryTreeNodes(leftDepth)
        rightTreeNodes = countNodesExec(root.Right)
    } else {
        leftTreeNodes = countNodesExec(root.Left)
        rightTreeNodes = getFullBinaryTreeNodes(rightDepth)
    }
    return leftTreeNodes + rightTreeNodes + 1
}

// 获取满二叉树的节点个数 (传入树的深度)
func getFullBinaryTreeNodes(depth int) int {
    return (1 << uint8(depth)) - 1
}

// 获取完全二叉树的深度 (每次都向左子树走)
func getCompleteBinaryTreeDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return getCompleteBinaryTreeDepth(root.Left) + 1
}
```
#####   3. 获取普通二叉树的深度
```go
// 求二叉树深度 (最大深度)
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left, right := maxDepth(root.Left), maxDepth(root.Right)
    return max(left, right) + 1
}

// 求二叉树深度 (最小深度)
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left, right := minDepth(root.Left), minDepth(root.Right)
    if left == 0 || right == 0 {
        return left + right + 1
    }
    return min(left, right) + 1
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
#####   4. 获取完全二叉树的深度
```go
func getCompleteBinaryTreeDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return getCompleteBinaryTreeDepth(root.Left) + 1
}   
```
##### 5. 获取二叉树左下角的值
```go
var maxDepth int // 最深层数
var ans int      // 最深层左下角的值

// 采用DFS找到二叉树左下角的值
func findBottomLeftValue(root *TreeNode) int {
    ans, maxDepth = -1, -1
    findBottomLeftValueExec(root, 0)
    return ans
}

func findBottomLeftValueExec(root *TreeNode, depth int) {
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        if depth > maxDepth {
            ans = root.Val
            maxDepth = depth
        }
    }
    findBottomLeftValueExec(root.Left, depth+1)
    findBottomLeftValueExec(root.Right, depth+1)
}
```
#####  6. 获取左叶子节点之和
```go
// 求二叉树左叶子节点之和		(解法1)
func sumOfLeftLeaves(root *TreeNode) int {
    return sumOfLeftLeavesExec(root)
}

func sumOfLeftLeavesExec(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        return root.Left.Val + sumOfLeftLeavesExec(root.Right)
    }
    return sumOfLeftLeavesExec(root.Left) + sumOfLeftLeavesExec(root.Right)
}

// 求二叉树左叶子节点之和		(解法2)
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return sumOfLeftLeavesExec(root.Left, 0) + sumOfLeftLeavesExec(root.Right, 1)
}

// flag == 0，表示此时在左子树。 (相对于父亲)
// flag == 1，表示此时在右子树。 (相对于父亲)
func sumOfLeftLeavesExec(root *TreeNode, flag int) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        if flag == 0 {
            return root.Val
        }
        return 0
    }
    return sumOfLeftLeavesExec(root.Left, 0) + sumOfLeftLeavesExec(root.Right, 1)
}
```

#### 性质判断
##### 1. 判断二叉树是否为二叉搜索树
```go
const inf = 1000000000000
var lastVal int // 当前中序遍历序列的最后一个数

// 判断二叉树是否为二叉搜索树
func isValidBST(root *TreeNode) bool {
    lastVal = -inf
    return isValidBSTExec(root)
}

func isValidBSTExec(root *TreeNode) bool {
    if root == nil{
        return true
    }
    if !isValidBSTExec(root.Left) || root.Val <= lastVal {
        return false
    }
    lastVal = root.Val
    return isValidBSTExec(root.Right)
}
```
##### 2. 判断二叉树是否为完全二叉树
```go
// 判断二叉树是否为完全二叉树
func isCompleteTree(root *TreeNode) bool {
    queue := []*TreeNode{root}
    nullAppearFlag := false // 这是一个标识位，标识遍历过程中是否出现了空节点
    for len(queue) != 0 {
        top := queue[0]
        queue = queue[1:]
        if top == nil {
            nullAppearFlag = true
            continue
        }
        if nullAppearFlag {
            return false
        }
        queue = append(queue, top.Left)
        queue = append(queue, top.Right)
    }
    return true
}
```
##### 3. 判断二叉树是否为单值二叉树
```go
// 判断二叉树是否为单值二叉树	(写法1)
func isUnivalTree(root *TreeNode) bool {
    return isUnivalTreeExec(root)
}

func isUnivalTreeExec(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left != nil && root.Left.Val != root.Val {
        return false
    }
    if root.Right != nil && root.Right.Val != root.Val {
        return false
    }
    return isUnivalTreeExec(root.Left) && isUnivalTreeExec(root.Right)
}

// 判断二叉树是否为单值二叉树	(写法2)
func isUnivalTree(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isUnivalTreeExec(root, root.Val)
}

func isUnivalTreeExec(root *TreeNode, number int) bool {
    if root == nil {
        return true
    }
    return root.Val == number && isUnivalTreeExec(root.Left, number) && isUnivalTreeExec(root.Right, number)
}
```
##### 4. 判断二叉树是否平衡
```go
// 判断二叉树是否为平衡二叉树	(写法1)
func isBalanced(root *TreeNode) bool {
    ans, _ := isBalancedExec(root)
    return ans
}

// 第一个返回值是该树是否平衡，第二个是该树的高度
func isBalancedExec(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }
    leftIsBalanced, leftHeight := isBalancedExec(root.Left)
    rightIsBalanced, rightHeight := isBalancedExec(root.Right)
    return leftIsBalanced && rightIsBalanced && abs(leftHeight-rightHeight) <= 1, max(leftHeight, rightHeight) + 1
}

// 判断二叉树是否为平衡二叉树	(写法2)
func isBalanced(root *TreeNode) bool {
    ans := isBalancedExec(root)
    return ans != -1
}

// 返回树高 (-1表示不是平衡二叉树)
func isBalancedExec(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHeight := isBalancedExec(root.Left)
    rightHeight := isBalancedExec(root.Right)
    if leftHeight != -1 && rightHeight != -1 && abs(leftHeight-rightHeight) <= 1 {
        return max(leftHeight, rightHeight) + 1
    }
    return -1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}
```
##### 5. 判断二叉树是否对称
```go
// 判断二叉树是否对称
func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}

// 判断两棵二叉树是否互为镜像
func isMirror(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}
```
#### 二叉树构建
##### 1. 通过先序、中序遍历序列构建二叉树
```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    return buildTreeExec(preorder, inorder)
}
func buildTreeExec(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    index, rootVal := 0, preorder[0]
    for i := 0; i < len(inorder); i++ {
        if inorder[i] == rootVal {
            index = i
            break
        }
    }
    root := &TreeNode{
        Val:   rootVal,
        Left:  buildTreeExec(preorder[1:index+1], inorder[:index]),
        Right: buildTreeExec(preorder[index+1:], inorder[index+1:]),
    }
    return root
}
```
##### 2. 通过后序、中序遍历序列构建二叉树
```go
func buildTree(inorder []int, postorder []int) *TreeNode {
    return buildTreeExec(inorder, postorder)
}
func buildTreeExec(inorder []int, postorder []int) *TreeNode {
    n := len(postorder)
    if n == 0 {
        return nil
    }
    index, rootVal := 0, postorder[n-1]
    for i := 0; i < n; i++ {
        if inorder[i] == rootVal {
            index = i
            break
        }
    }
    root := &TreeNode{
        Val:   rootVal,
        Left:  buildTreeExec(inorder[:index], postorder[:index]),
        Right: buildTreeExec(inorder[index+1:], postorder[index:n-1]),
    }
    return root
}
```
##### 3. 通过先序遍历序列构造二叉搜索树
```go
func bstFromPreorder(preorder []int) *TreeNode {
    return bstFromPreorderExec(preorder)
}
func bstFromPreorderExec(preorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    index, rootVal := 0, preorder[0]
    for index = 0; index < len(preorder); index++ {
        if preorder[index] > rootVal {
            break
        }
    }
    root := &TreeNode{
        Val:   rootVal,
        Left:  bstFromPreorderExec(preorder[1:index]),
        Right: bstFromPreorderExec(preorder[index:]),
    }
    return root
}
```
##### 4. 通过中序遍历序列构造平衡二叉搜索树
```go
func sortedArrayToBST(nums []int) *TreeNode {
    return sortedArrayToBSTExec(nums)
}
func sortedArrayToBSTExec(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums) >> 1
    root := &TreeNode{
        Val:   nums[mid],
        Left:  sortedArrayToBSTExec(nums[:mid]),
        Right: sortedArrayToBSTExec(nums[mid+1:]),
    }
    return root
}
```
#### 二叉树结构调整
##### 1. 翻转二叉树
```go
// 翻转二叉树   (递归写法)
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Left, root.Right = root.Right, root.Left
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}

// 翻转二叉树   (迭代写法)
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) != 0 {
        top := queue[0]
        queue = queue[1:]
        top.Left, top.Right = top.Right, top.Left
        if top.Left != nil {
            queue = append(queue, top.Left)
        }
        if top.Right != nil {
            queue = append(queue, top.Right)
        }
    }
    return root
}
```
##### 2. 删除二叉树不含 `1` 的子树
```go
// 删除二叉树不含1的子树
func pruneTree(root *TreeNode) *TreeNode {
    return pruneTreeExec(root)
}
func pruneTreeExec(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Left = pruneTreeExec(root.Left)
    root.Right = pruneTreeExec(root.Right)
    if root.Left == nil && root.Right == nil && root.Val != 1 {
        return nil
    }
    return root
}
```
##### 3. 二叉树展开为链表
```go
var headNode *TreeNode // 链表的头节点

// 反后序遍历将二叉树展开为链表
func flatten(root *TreeNode) {
    headNode = nil
    flattenExec(root)
}
func flattenExec(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    root.Right = flattenExec(root.Right)
    root.Left = flattenExec(root.Left)

    root.Right = headNode
    root.Left = nil
    headNode = root
    return root
}
```
##### 4. 合并二叉树 
[传送门](https://leetcode-cn.com/problems/merge-two-binary-trees/)
```go
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    return mergeTreesExec(t1, t2)
}
func mergeTreesExec(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }
    if t1 == nil {
        return t2
    }
    if t2 == nil {
        return t1
    }
    t1.Left = mergeTreesExec(t1.Left, t2.Left)
    t1.Right = mergeTreesExec(t1.Right, t2.Right)
    t1.Val = t1.Val + t2.Val
    return t1
}
```
##### 5. 将二叉搜索树转换为累加树 
[传送门](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/comments/)
```go
var lastSum int    // 保存当前遍历过的节点值总和
func convertBST(root *TreeNode) *TreeNode {
    lastSum = 0
    convertBSTExec(root)
    return root
}
func convertBSTExec(root *TreeNode) {
    if root == nil {
        return
    }
    convertBSTExec(root.Right)
    root.Val += lastSum
    lastSum = root.Val
    convertBSTExec(root.Left)
}
```
##### 6. 删除二叉树值为目标值的节点，返回删除后形成的森林
[传送门](https://leetcode-cn.com/problems/delete-nodes-and-return-forest/submissions/)
```go
var forest []*TreeNode
var shouldDelete map[int]bool
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    forest = make([]*TreeNode, 0)
    shouldDelete = make(map[int]bool)
    for i := 0; i < len(to_delete); i++ {
        shouldDelete[to_delete[i]] = true
    }
    delNodesExec(root, true)
    return forest
}

// shouldAddFlag: 是一个标识，标识该树是否需要加入结果集
// 这个函数实现了2个功能: (1) 调整二叉树结构  (2) 形成结果集
// 其中 if shouldAddFlag && root != nil {...} 就是为了形成结果集，其它的代码就是为了调整二叉树结构
func delNodesExec(root *TreeNode, shouldAddFlag bool) *TreeNode {
    if root == nil {
        return root
    }
    deleteFlag := shouldDelete[root.Val]

    // 如果根节点在to_delete中，那么我们需要把它的左右子树加入结果集。
    //          (前提是左右子树非空，这也是为什么形成结果集那里加了一个 root != nil 的判断)。
    // 如果根节点没在to_delete中，那么我们不能把它的左右子树加入结果集，而是把该根节点加入结果集。
    root.Left = delNodesExec(root.Left, deleteFlag)
    root.Right = delNodesExec(root.Right, deleteFlag)
    if deleteFlag {
        root = nil
    }
    // 形成结果集(森林)
    if shouldAddFlag && root != nil {
        forest = append(forest, root)
    }
    return root
}
```

####  二叉树遍历 (返回遍历序列)
#####  1. 先序遍历
```go
// 递归方式
var preorderSequence []int
func preorderTraversal(root *TreeNode) []int {
    preorderSequence = make([]int, 0)
    preorderTraversalExec(root)
    return preorderSequence
}
func preorderTraversalExec(root *TreeNode) {
    if root == nil {
        return
    }
    preorderSequence = append(preorderSequence, root.Val)
    preorderTraversalExec(root.Left)
    preorderTraversalExec(root.Right)
}

// 迭代方式
func preorderTraversal(root *TreeNode) []int {
    preorderSequence := make([]int, 0)
    isVisited := make(map[*TreeNode]int)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    isVisited[root] = 0
    for len(stack) != 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if top == nil {
            continue
        }
        if isVisited[top] == 0 {
            stack = append(stack, top.Right)
            stack = append(stack, top.Left)
            stack = append(stack, top)
            isVisited[top.Right] = 0
            isVisited[top.Left] = 0
            isVisited[top] = 1
        } else {
            preorderSequence = append(preorderSequence, top.Val)
        }
    }
    return preorderSequence
}
```
##### 2. 中序遍历
```go
// 递归方式
var inorderSequence []int
func inorderTraversal(root *TreeNode) []int {
    inorderSequence = make([]int,0)
    inorderTraversalExec(root)
    return inorderSequence
}
func inorderTraversalExec(root *TreeNode) {
    if root == nil{
        return
    }
    inorderTraversalExec(root.Left)
    inorderSequence = append(inorderSequence,root.Val)	// 与先序遍历相比，就这语句调整了下位置。
    inorderTraversalExec(root.Right)
}

// 迭代方式
func inorderTraversal(root *TreeNode) []int {
    inorderSequence := make([]int, 0)
    isVisited := make(map[*TreeNode]int)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    isVisited[root] = 0
    for len(stack) != 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if top == nil {
            continue
        }
        if isVisited[top] == 0 {
            stack = append(stack, top.Right)
            stack = append(stack, top) // 与先序遍历迭代相比，也只是调整了这句语句的位置
            stack = append(stack, top.Left)
            isVisited[top.Right] = 0
            isVisited[top] = 1
            isVisited[top.Left] = 0
        } else {
            inorderSequence = append(inorderSequence, top.Val)
        }
    }
    return inorderSequence
}
```
##### 3. 后序遍历
```go
// 递归方式
var postorderSequence []int
func postorderTraversal(root *TreeNode) []int {
    postorderSequence = make([]int, 0)
    postorderTraversalExec(root)
    return postorderSequence
}
func postorderTraversalExec(root *TreeNode) {
    if root == nil {
        return
    }
    postorderTraversalExec(root.Left)
    postorderTraversalExec(root.Right)
    postorderSequence = append(postorderSequence, root.Val)
}

// 迭代方式
func postorderTraversal(root *TreeNode) []int {
    postorderSequence := make([]int, 0)
    isVisited := make(map[*TreeNode]int)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    isVisited[root] = 0
    for len(stack) != 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if top == nil {
            continue
        }
        if isVisited[top] == 0 {
            stack = append(stack, top)
            stack = append(stack, top.Right)
            stack = append(stack, top.Left)
            isVisited[top.Right] = 0
            isVisited[top.Left] = 0
            isVisited[top] = 1
        } else {
            postorderSequence = append(postorderSequence, top.Val)
        }
    }
    return postorderSequence
}
```
##### 4. 层序遍历
```go
// 递归方式
var levelOrderSequence [][]int
func levelOrder(root *TreeNode) [][]int {
    levelOrderSequence = make([][]int, 0)
    levelOrderExec(root, 0)
    return levelOrderSequence
}
func levelOrderExec(root *TreeNode, lay int) {
    // lay 表示当前树节点所在的层 (从0开始)
    if root == nil {
        return
    }
    if len(levelOrderSequence) == lay {
        levelOrderSequence = append(levelOrderSequence, []int{})
    }
    levelOrderSequence[lay] = append(levelOrderSequence[lay], root.Val)
    levelOrderExec(root.Left, lay+1)
    levelOrderExec(root.Right, lay+1)
}

// 迭代方式
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    levelOrderSequence := make([][]int, 0)
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    for len(queue) != 0 {
        size := len(queue)
        nowOrderSequence := make([]int, 0)
        for i := 0; i < size; i++ {
            top := queue[0]
            queue = queue[1:]
            nowOrderSequence = append(nowOrderSequence, top.Val)
            if top.Left != nil {
                queue = append(queue, top.Left)
            }
            if top.Right != nil {
                queue = append(queue, top.Right)
            }
        }
        levelOrderSequence = append(levelOrderSequence, nowOrderSequence)
    }
    return levelOrderSequence
}
```
#### 二叉树间关系判断
##### 1. 判断是否对称、互为镜像
```go
// 判断二叉树是否对称
func isSymmetric(root *TreeNode) bool {
    return isMirror(root, root)
}

// 判断两棵二叉树树是否互为镜像
func isMirror(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}
```
##### 2. 判断是否相同
```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```
### 练习题
1. 属性获取
    - [ ] [104. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
    - [ ] [111. 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)
    - [ ] [222. 完全二叉树的节点个数](https://leetcode-cn.com/problems/count-complete-tree-nodes/)
    - [ ] [404. 左叶子之和](https://leetcode-cn.com/problems/sum-of-left-leaves/)
    - [ ] [513. 找树左下角的值](https://leetcode-cn.com/problems/find-bottom-left-tree-value/submissions/)
    - [ ] [563. 二叉树的坡度](https://leetcode-cn.com/problems/binary-tree-tilt/)
2. 性质判断
    - [ ] [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/submissions)
    - [ ] [101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/submissions/)
    - [ ] [110. 平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/)
    - [ ] [958. 二叉树完全性检验](https://leetcode-cn.com/problems/check-completeness-of-a-binary-tree/)
    - [ ] [965. 单值二叉树](https://leetcode-cn.com/problems/univalued-binary-tree/)

3. 二叉树构建
    - [ ] [105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/submissions/)
    - [ ] [106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
    - [ ] [108. 将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/submissions/)
    - [ ] [654. 构造最大二叉树](https://leetcode-cn.com/problems/maximum-binary-tree/)
    - [ ] [1008. 先序遍历构造二叉树](https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/submissions/)

4. 二叉树结构调整
    - [ ] [114. 二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/submissions/)
    - [ ] [226. 翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/submissions/)
    - [ ] [538. 把二叉搜索树转换为累加树](https://leetcode-cn.com/problems/convert-bst-to-greater-tree/comments/)
    - [ ] [617. 合并二叉树](https://leetcode-cn.com/problems/merge-two-binary-trees/)
    - [ ] [814. 二叉树剪枝](https://leetcode-cn.com/problems/binary-tree-pruning/)
    - [ ] [1110. 删点成林](https://leetcode-cn.com/problems/delete-nodes-and-return-forest/submissions/)
5. 二叉树遍历
    - [ ] [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/submissions/)
    - [ ] [102. 二叉树的层次遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
    - [ ] [103. 二叉树的锯齿形层次遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/)
    - [ ] [107. 二叉树的层次遍历 II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/)
    - [ ] [144. 二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)
    - [ ] [145. 二叉树的后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/submissions/)
    
6. 二叉树间关系判断
    - [ ] [100. 相同的树](https://leetcode-cn.com/problems/same-tree/submissions/)
    - [ ] [101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/submissions/)
