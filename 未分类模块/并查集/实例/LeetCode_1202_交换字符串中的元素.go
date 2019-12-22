package main

import "sort"

var father []int
// 并查集初始化
func initUFS(size int) {
	father = make([]int, size)
	for i := 0; i < size; i++ {
		father[i] = i
	}
}

// 合并两个并查集
func mergeUFS(a, b int) {
	// 注意a,b必须 < father数组的长度
	father1, father2 := findFather(a), findFather(b)
	father[father1] = father2 // 这样也可以
	/*
		father[father1] = min(father1,father2)   // 这种写法也可以，可以保证大的指向小的
		father[father2] = min(father1,father2)   // 这种写法也可以，可以保证大的指向小的
	*/
}

// 路径压缩的并查集找爸爸
func findFather(i int) int {
	if i == father[i] {
		return i
	}
	father[i] = findFather(father[i]) // 路径压缩
	return father[i]
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	belongList := make(map[int][]uint8)

	initUFS(len(s))	// 初始化并查集
	// 合并
	for i := 0; i < len(pairs); i++ {
		mergeUFS(pairs[i][0], pairs[i][1])
	}
	// 将s[i]放入集合leader所对应的切片 (leader在这就是下面的fatherI)
	for i := 0; i < len(s); i++ {
		fatherI := findFather(i)
		belongList[fatherI] = append(belongList[fatherI], s[i])
	}
	// 按字符大小排序
	for _, list := range belongList {
		sort.Slice(list, func(i, j int) bool {
			return list[i] < list[j]
		})
	}

	// 写回原字符串
	bytes := []byte(s)
	for i := 0; i < len(s); i++ {
		fatherI := findFather(i)
		bytes[i] = belongList[fatherI][0]				// 这一句是取队首	(由于go没有队列，所以只能有切片模拟队列了)
		belongList[fatherI] = belongList[fatherI][1:]	// 这一句是出队		(由于go没有队列，所以只能有切片模拟队列了)
	}
	return string(bytes)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return b
}
/*

*/


/*
	总结
	1. 这题的总体思路是:
			(1) 让可以相互交换字符的索引组成一个集合，这样就形成了多个集合。							(构建集合通过并查集实现)
			(2) 将每个集合中索引对应的字符加入到对应的切片中，对切片进行排序。							(可以让每个集合有一个leader,然后将该集合的所有字符放入leader所对应的切片)
																										(上面代码的leader就是集合的根节点)
			(3) 遍历索引[0,len(s)-1],记为i，将该索引i对应的字符写回s[i]中。								(由于go的string不能修改，所以我采用了[]byte实现)
*/