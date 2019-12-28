package main

/*
	给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。
	对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。
	你可以返回任何满足上述条件的数组作为答案。
*/

func sortArrayByParityII(A []int) []int {
	writer, reader := 0, 0 // 这里使writer指向第一个不满足要求的位置，即writer之前的位置都是满足要求的
	for reader < len(A) {
		for writer < len(A) && A[writer]%2 == writer%2 {
			writer++
		}
		reader = writer // reader找满足writer位置要求的数
		for reader < len(A) && A[reader]%2 == A[writer]%2 {
			reader++
		}
		if reader < len(A) {
			A[writer], A[reader] = A[reader], A[writer]
		}
	}
	return A
}

/*
	题目链接:
		https://leetcode-cn.com/problems/sort-array-by-parity-ii				按奇偶排序数组2
*/


/*
	总结
	1. 这题目如果不要求原地修改的话，那可以把奇偶数遍历时放在另外的数组。(
			记录2个指针，一个指向奇数位置，一个指向偶数位置，初始时为(1,0)。
			之后读取原数组中的数据。每次读取一个奇数时，写入奇数指针的位置后奇数指针+=2，
			如果是偶数则写入偶数指针对应的位置后，偶数指针+=2。
		)
		之后再输出该数组就可以了。
	2. 以上采用了双指针的方法实现了原地修改。(虽然感觉这个双指针可能有点勉强)
*/
