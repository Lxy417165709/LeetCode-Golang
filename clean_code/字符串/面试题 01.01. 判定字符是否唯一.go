package 字符串

import (
	"bytes"
	"sort"
)

// -------------- 不使用任何数据结构的版本 --------------
func isUnique(astr string) bool {
	for i:=0;i<len(astr);i++{
		for t:=i+1;t<len(astr);t++{
			if astr[i]==astr[t]{
				return false
			}
		}
	}
	return true
}

func isUnique(astr string) bool {
	for i:=0;i<len(astr);i++{
		if bytes.LastIndex([]byte(astr),[]byte{astr[i]})!=bytes.Index([]byte(astr),[]byte{astr[i]}){
			return false
		}
	}
	return true
}

func isUnique(astr string) bool {
	bytes := []byte(astr)
	sort.Slice(bytes,func(i,j int) bool{
		return bytes[i]<bytes[j]
	})
	for i:=1;i<len(bytes);i++{
		if bytes[i]==bytes[i-1]{
			return false
		}
	}
	return true
}

func isUnique(astr string) bool{
	// 这个也相当于哈希，只不过是采用类似Bitmap的方式存储哈希映射关系。
	hashArray := 0
	for i:=0;i<len(astr);i++{
		bitValueOfChar := 1<<(astr[i]-'a')
		if hashArray & bitValueOfChar != 0{
			return false
		}
		hashArray |= bitValueOfChar
	}
	return true
}

// -------------- 使用哈希的版本 --------------
func isUnique(astr string) bool {
	countOfChar := make(map[byte]int)
	for i:=0;i<len(astr);i++{
		countOfChar[astr[i]]++
	}
	for _,count := range countOfChar{
		if count>1{
			return false
		}
	}
	return true
}


/*
	题目链接: https://leetcode-cn.com/problems/is-unique-lcci/submissions/
	总结:
		1. 不使用数据结构，还可以使用find函数、排序解决这个题目。
		2. 官方有使用位运算作为底层哈希数组，这也没用到任何数据结构。 (相当于bit map)
*/
