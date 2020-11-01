package main

var ans []string

func letterCasePermutation(S string) []string {
	ans = make([]string,0)
	letterCasePermutationExec([]byte(S),[]byte{})
	return ans
}

func letterCasePermutationExec(bytes []byte,seq []byte) {
	if len(bytes)==0{
		ans = append(ans,string(seq))
		return
	}
	if bytes[0]>='a' && bytes[0]<='z' {
		letterCasePermutationExec(bytes[1:],append(seq,bytes[0]-'a'+'A'))
	}
	if bytes[0]>='A' && bytes[0]<='Z' {
		letterCasePermutationExec(bytes[1:],append(seq,bytes[0]-'A'+'a'))
	}
	letterCasePermutationExec(bytes[1:],append(seq,bytes[0]))
}

/*
	总结
	1. 这题刚开始做的时候，在letterCasePermutationExec函数中采用了一个for循环遍历bytes，然后发现是错误的。
		理清思路后，去除了for循环，然后就AC了。
*/