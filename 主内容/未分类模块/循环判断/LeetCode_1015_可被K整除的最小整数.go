package 循环判断

// 哈希判断循环
func smallestRepunitDivByK(K int) int {
	length := 0 // N的长度
	curNum := 0 // 当前的N % K
	isVisit := make(map[int]bool)
	for isVisit[curNum] == false {
		isVisit[curNum] = true
		curNum = curNum*10 + 1
		length ++
		curNum %= K
	}
	if curNum == 0 {
		return length
	}
	return -1
}

// 快慢指针判断循环
func smallestRepunitDivByK(K int) int {
	length := 0 // N的长度
	slowNum, fastNum := 0, 0
	for true {
		length++
		slowNum = slowNum*10 + 1
		fastNum = fastNum*10 + 1
		fastNum = fastNum*10 + 1
		slowNum %= K
		fastNum %= K
		if slowNum == fastNum {
			break
		}
	}
	if slowNum == 0 {
		return length
	}
	return -1
}

/*
    总结
    1. 这题目没有使用数学方法，而是利用哈希判断是否存在循环，如果存在循环，且该循环点不为0，那么就不可能
        存在N，否则输出此时的长度。
    2. 这题是看了官方大佬的题解之后写出来的，想法还不太清晰。
    3. 除了使用哈希判断是否存在循环外，还可以使用快慢指针喔。
	4. 或许对于这类数论题来说，判断循环是一个好方法！
*/
