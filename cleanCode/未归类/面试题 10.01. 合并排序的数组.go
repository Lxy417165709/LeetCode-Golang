package 未归类

func merge(A []int, m int, B []int, n int) {
	posOfA := m - 1
	posOfB := n - 1
	posOfWriting := m + n - 1
	for posOfB >= 0 && posOfA >= 0 {
		if A[posOfA] > B[posOfB] {
			A[posOfWriting] = A[posOfA]
			posOfA--
		} else {
			A[posOfWriting] = B[posOfB]
			posOfB--
		}
		posOfWriting--
	}
	for posOfB >= 0 {
		A[posOfWriting] = B[posOfB]
		posOfWriting--
		posOfB--
	}
}
