package main

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	slow, fast := n, n
	for {
		slow = getHappy(slow)
		fast = getHappy(getHappy(fast))
		if slow == 1 {
			return true
		}
		if slow == fast {
			return false
		}

	}

}

func getHappy(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

// 通过快慢指针判环就可以了
