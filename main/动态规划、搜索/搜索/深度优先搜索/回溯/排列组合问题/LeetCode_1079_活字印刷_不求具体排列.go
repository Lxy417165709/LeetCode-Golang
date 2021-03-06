package 排列组合问题

// 这题目就是要求: 有重复元素的数组的所有排列。 这个代码不求具体的，而是算有多少个不同的全排列。
func numTilePossibilities(tiles string) int {
	ans = 0
	numTilePossibilitiesExec([]byte(tiles), 0)
	return ans - 1 // -1 是为了去除空的
}

var ans int
func numTilePossibilitiesExec(bytes []byte, index int) {

	/*
		if index == len(bytes){
			ans++
		}
		// 如果代码是这样的话，那么求的是bytes的全排列种数，而如果是直接ans++
		// 那么求的
	*/
	ans++	// 总结点①
	isVisit := make(map[uint8]bool)
	// 获取 [index,len(bytes)-1]的全排列
	for i := index; i < len(bytes); i++ {
		if isVisit[bytes[i]] == true {
			continue
		}
		isVisit[bytes[i]] = true
		bytes[index], bytes[i] = bytes[i], bytes[index]
		numTilePossibilitiesExec(bytes, index+1)
		bytes[index], bytes[i] = bytes[i], bytes[index]
	}
}
/*
	总结
	1. 总结点①:
		// 如果代码是下面这样的话，那么求的是 bytes 的全排列种数。
		if index == len(bytes){
			ans++
		}

		// 如果代码直接是 ans++，那么求的是 bytes 以及它的子集的全排列种数。

		即，通过限制，我们可以获得 bytes 长度在 [0,len(bytes)-1] 的全排列种数。
*/
