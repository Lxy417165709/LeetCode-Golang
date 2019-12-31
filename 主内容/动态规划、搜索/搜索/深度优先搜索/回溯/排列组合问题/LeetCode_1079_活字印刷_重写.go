package 排列组合问题

// 这题目就是要求: 有重复元素的数组的所有排列。
func numTilePossibilities(tiles string) int {
	ans = make([]string, 0)
	numTilePossibilitiesExec([]byte(tiles), 0,[]byte{})
	return len(ans) - 1 // -1 是为了去除空的
}

var ans []string
func numTilePossibilitiesExec(bytes []byte, index int,sequence []byte) {

	ans = append(ans, string(sequence))
	isVisit := make(map[uint8]bool)
	// 获取 [index,len(bytes)-1]的全排列
	for i := index; i < len(bytes); i++ {
		if isVisit[bytes[i]] == true {
			continue
		}
		isVisit[bytes[i]] = true
		bytes[index], bytes[i] = bytes[i], bytes[index]
		numTilePossibilitiesExec(bytes, index+1,append(sequence, bytes[index]))
		bytes[index], bytes[i] = bytes[i], bytes[index]
	}
}
