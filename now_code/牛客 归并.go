package main

func merge( A []int ,  m int, B []int, n int )  {
	a := m-1
	b := n-1
	curStoreIndex := m+n-1
	for a>=0 && b>=0{
		if A[a]>=B[b]{
			A[curStoreIndex]=A[a]
			a--
		}else{
			A[curStoreIndex]=B[b]
			b--
		}
		curStoreIndex--
	}
	for i:=0;i<=a;i++{
		A[i]=A[i]
	}
	for i:=0;i<=b;i++{
		A[i]=B[i]
	}
}

/*
	题目链接: https://www.nowcoder.com/practice/89865d4375634fc484f3a24b7fe65665?tpId=188&&tqId=36160&rp=1&ru=/ta/job-code-high-week&qru=/ta/job-code-high-week/question-ranking
*/