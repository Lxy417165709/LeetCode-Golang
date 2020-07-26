package 前缀和

func findLongestSubarray(array []string) []string {
	indexOfDiffCountOfLetterAndDigit := make(map[int]int)
	indexOfDiffCountOfLetterAndDigit[0] = -1
	diffCountOfLetterAndDigit := 0
	longestSubarray := make([]string, 0)
	for i := 0; i < len(array); i++ {
		if isLetter(array[i][0]) {
			diffCountOfLetterAndDigit++
		} else {
			diffCountOfLetterAndDigit--
		}
		if index, ok := indexOfDiffCountOfLetterAndDigit[diffCountOfLetterAndDigit]; !ok {
			indexOfDiffCountOfLetterAndDigit[diffCountOfLetterAndDigit] = i
		} else {
			nowSubarrayLength := i - index
			longestSubarrayLength := len(longestSubarray)
			if longestSubarrayLength < nowSubarrayLength {
				longestSubarray = array[i+1-nowSubarrayLength : i+1]
			}
		}
	}
	return longestSubarray
}

func isLetter(char byte) bool{
	return  char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z'
}


/*
	题目链接: https://leetcode-cn.com/problems/find-longest-subarray-lcci/
	总结
		1. 一开始想到使用前缀和了，但是不知道怎么解，之后想到滑动窗口，动态规划，都没解出来
		2. 看了解答后，理清思路后就直接AC了。
		3. 要注意这题的字母可能是大写字母、小写字母
		4. 思路就是 将字母当做+1，数字当做-1，扫描字符串，利用哈希记录字母(+1)、数字(-1)的差值与索引的关系， (依靠题目条件，只记录最小的索引)
			则下次遇到相同的差值时，找到哈希表中对应的先前索引，利用运算就能知道最长的串了。
		5. 这题和 525. 连续数组 一样。
*/
