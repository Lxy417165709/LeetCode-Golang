package 排列组合问题

// 这题目就是要求: 有重复元素的数组的所有排列。
func numTilePossibilities(tiles string) int {
	ans = make([]string, 0)
	numTilePossibilitiesExec([]byte(tiles), []byte{})
	return len(ans) - 1 // -1 是为了去除空的
}

var ans []string

func numTilePossibilitiesExec(bytes []byte, sequence []byte) {
	ans = append(ans, string(sequence))
	isVisit := make(map[uint8]bool)
	for i := 0; i < len(bytes); i++ {
		if isVisit[bytes[i]] == true {
			continue
		}
		isVisit[bytes[i]] = true
		bytes[0], bytes[i] = bytes[i], bytes[0]
		numTilePossibilitiesExec(bytes[1:], append(sequence, bytes[0]))
		bytes[0], bytes[i] = bytes[i], bytes[0]
	}
}
