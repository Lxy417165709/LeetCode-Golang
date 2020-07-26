package 动态规划

func fib(n int) int {
	fibonacciSequence := NewFibonacciSequence(100, 1000000007)
	return fibonacciSequence.GetNumber(n)
}

// ----------- FibonacciSequence -----------
type FibonacciSequence struct {
	numbers       []int
	maxVisitIndex int
	mod           int
}

func NewFibonacciSequence(maxVisitIndex int, mod int) *FibonacciSequence {
	fibonacciSequence := &FibonacciSequence{maxVisitIndex: maxVisitIndex, mod: mod}
	fibonacciSequence.formNumbers()
	return fibonacciSequence
}

func (fs *FibonacciSequence) GetNumber(index int) int {
	return fs.numbers[index]
}

func (fs *FibonacciSequence) formNumbers() {
	fs.numbers = make([]int, fs.maxVisitIndex+1)
	fs.numbers[0], fs.numbers[1] = 0, 1
	for i := 2; i < len(fs.numbers); i++ {
		fs.numbers[i] = (fs.numbers[i-1] + fs.numbers[i-2]) % fs.mod
	}
}
