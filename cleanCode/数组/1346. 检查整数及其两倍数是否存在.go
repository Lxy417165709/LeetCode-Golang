package 数组
// ----------------------- 哈希 -----------------------
func checkIfExist(arr []int) bool {
	isExist := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if isExist[arr[i]*2] || arr[i]%2 == 0 && isExist[arr[i]/2] {
			return true
		}
		isExist[arr[i]] = true
	}
	return false
}
/*
	题目链接: https://leetcode-cn.com/problems/check-if-n-and-its-double-exist/
*/
