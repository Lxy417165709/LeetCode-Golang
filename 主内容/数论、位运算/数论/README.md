# 数论 (TODO)

## 1. 质数
### 1.1 判断 `整数x` 是否为质数
```go
func isPrime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
```

### 1.2 获取 `[1, n]` 中的质数 (素数筛)
```go
func primeGenerator(n int) []int {
	isNotPrime := make([]bool, n+1)
	ans := []int{}

	for i := 2; i*i <= n; i++ {
		if isNotPrime[i] == true {
			continue
		}
		for t := i * i; t <= n; t += i {
			isNotPrime[t] = true
		}
	}
	for i := 2; i <= n; i++ {
		if isNotPrime[i] == false {
			ans = append(ans, i)
		}
	}
	return ans
}
```

## 2. 因子
### 2.1 求 `整数x` 的因子和
```go
func getFactorSum(x int) int {
	sum := 0
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			sum += i
			if x/i != i {
				sum += x / i
			}
		}
	}
	return sum
}
```

### 2.2 求 `整数x` 与 `整数y` 最大公约数
```go
func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%b)
}
```

### 2.3 求 `n!` 尾部 `0` 的个数
```go
func trailingZeroes(n int) int {
	ans := 0
	for n != 0 {
		ans += n / 5
		n = n / 5
	}
	return ans
}
```
### 2.4 判断 `整数x` 是否只有某些因子
```go
// 判断 x 是否只有 factorSet 中的因子
func isOnlyHaving(factorSet []int, x int) bool {
	for i := 0; i < len(factorSet); i++ {
		for x%factorSet[i] == 0 {
			x /= factorSet[i]
		}
	}
	return x == 1
}
```

## 3. 定理
### 3.1 裴蜀定理
> $设a_{1},a_{2},a_{3}......a_{n}为n个整数，d是它们的最大公约数，那么存在整数x_{1}......x_{n}使得x_{1}*a_{1}+x_{2}*a_{2}+...x_{n}*a_{n}=d。$<br>
> [描述出处](https://baike.baidu.com/item/%E8%A3%B4%E8%9C%80%E5%AE%9A%E7%90%86/5186593?fr=aladdin)

[定理应用](https://leetcode-cn.com/problems/check-if-it-is-a-good-array/comments/)

### 3.2 费马小定理
> $如果p是一个质数，而整数a不是p的倍数，则有 a^{(p-1)}≡1（mod p）$。 <br>
> [描述出处](https://baike.baidu.com/item/%E8%B4%B9%E9%A9%AC%E5%B0%8F%E5%AE%9A%E7%90%86/4776158?fr=aladdin)

## 4. 练习题
- [ ] [172. 阶乘后的零](https://leetcode-cn.com/problems/factorial-trailing-zeroes/comments/)
- [ ] [204. 计数质数](https://leetcode-cn.com/problems/count-primes/)
- [ ] [263. 丑数](https://leetcode-cn.com/problems/ugly-number/submissions/)
- [ ] [507. 完美数](https://leetcode-cn.com/problems/perfect-number/submissions)
- [ ] [1250. 检查「好数组」](https://leetcode-cn.com/problems/check-if-it-is-a-good-array/comments/)

