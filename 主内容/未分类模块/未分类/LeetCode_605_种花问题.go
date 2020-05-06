package main

func canPlaceFlowers(flowerbed []int, n int) bool {

	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			n--
			flowerbed[i] = 1
		}
	}
	return n <= 0
}

// 总结：首尾+0，统一判断， 遇到3个0在一起时，就在中间插花。
