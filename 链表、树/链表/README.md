@[toc]

## 链表小总结
### 代码
#### 属性获取
##### 1. 获取链表环入口
```go
// 快慢指针法获取环入口
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head
    // 判断链表是否有环
    loopFlag := false
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            loopFlag = true
            break
        }
    }
    // 表示链表无环
    if loopFlag == false {
        return nil
    }
    // 表示链表有环
    loopEntrance := head
    for loopEntrance != slow {
        loopEntrance = loopEntrance.Next
        slow = slow.Next
    }
    return loopEntrance
}

```
##### 2. 获取两个链表相交节点
```go
// 优雅解法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pA, pB := headA, headB
    for pA != pB {
        if pA == nil {
            pA = headB
        } else {
            pA = pA.Next
        }

        if pB == nil {
            pB = headA
        } else {
            pB = pB.Next
        }
    }
    return pA
}
```
##### 3. 获取链表中点
```go
// 双指针法
func middleNode(head *ListNode) *ListNode {
    if head == nil{
       return nil
    }
    slow, fast := head, head
    //  slow, fast := head, head		如果这样写的话，链表节点为偶数时，返回: 偏右的中间节点。
    //  slow, fast := head, head.Next	如果这样写的话，链表节点为偶数时，返回: 偏左的中间节点。
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
```
#### 结构调整
##### 1. 反转链表
```go
// 迭代解法
func reverseList(head *ListNode) *ListNode {
    var pre *ListNode = nil
    var next *ListNode = nil
    for head != nil {
        next = head.Next
        head.Next = pre
        pre = head
        head = next
    }
    return pre
}

// 递归解法
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
```
##### 2. 删除链表倒数第n个节点
```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	// 定义哑头节点，这样可以方便删除头节点时的处理
	dummyHead := &ListNode{-1, head}
	left, right := dummyHead, dummyHead
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	// 此时left指向了要删除节点的父节点
	left.Next = left.Next.Next
	return dummyHead.Next
}
```


#### 性质判断
##### 1. 判断链表是否有环
```go
// 哈希解法
func hasCycle(head *ListNode) bool {
    isVisited := make(map[*ListNode]int)
    for head != nil {
        if _, ok := isVisited[head]; ok {
            return true
        }
        isVisited[head] = 1
        head = head.Next
    }
    return false
}

// 双指针解法
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

### 练习题
1. 属性获取
 	- [ ] [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)
 	- [ ] [160. 相交链表](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/)
 	- [ ] [876. 链表的中间结点](https://leetcode-cn.com/problems/middle-of-the-linked-list/)
3. 性质判断
  	- [ ] [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)
4. 结构调整
    - [ ] [19. 删除链表的倒数第N个节点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/comments/)
	- [ ] [92. 反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/)
  	- [ ] [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)
    

    
