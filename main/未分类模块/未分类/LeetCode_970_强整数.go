package main

import "math"

// low方法
func powerfulIntegers(x int, y int, bound int) []int {
	m := make(map[int]int)
	for i := 0; i < 20; i++ {
		a := int(math.Pow(float64(x), float64(i)))
		if a >= bound {
			break
		}
		for t := 0; t < 20; t++ {
			b := int(math.Pow(float64(y), float64(t)))
			if a+b > bound {
				break
			}
			m[a+b] = 1
		}
	}
	result := make([]int, 0)
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

// 更好的写法
func powerfulIntegers(x int, y int, bound int) []int {
	m := make(map[int]int)
	first, second := 1, 1
	for i := 0; i < 20 && first <= bound; i++ {
		second = 1
		for t := 0; t < 20 && first+second <= bound; t++ {
			m[first+second] = 1
			second *= y
		}
		first *= x
	}
	result := make([]int, 0)
	for k := range m {
		result = append(result, k)
	}
	return result
}

// 总结：当数字取幂时，一般可以采用穷举法
