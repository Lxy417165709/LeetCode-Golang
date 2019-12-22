package main

var father []int
// 初始化并查集
func initUFS(length int) {
	father = make([]int, length)
	for i := 0; i < length; i++ {
		father[i] = i
	}
}

// 合并并查集
func unionUFS(a, b int) {
	father1, father2 := findFather(a), findFather(b)
	father[father1] = father2
}

// 找爸爸
func findFather(a int) int {
	if a == father[a] {
		return a
	}
	father[a] = findFather(father[a])
	return father[a]
}

func regionsBySlashes(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	length := len(grid)
	initUFS(length * length * 4)

	for i := 0; i < length; i++ {
		for t := 0; t < length; t++ {
			start := 4 * (i*length + t)
			switch grid[i][t] {
			case ' ':
				unionUFS(start, start+1)
				unionUFS(start, start+2)
				unionUFS(start, start+3)
			case '\\':
				unionUFS(start, start+1)
				unionUFS(start+2, start+3)
			case '/':
				unionUFS(start, start+3)
				unionUFS(start+1, start+2)
			}
			if i > 0 {
				unionUFS(start, start-length*4+2)
			}
			if t > 0 {
				unionUFS(start-3, start+3)
			}
		}
	}
	ans := 0
	for i := 0; i < length*length*4; i++ {
		if i == findFather(i) {
			ans++
		}
	}
	return ans

}

/*
    总结
    1. 这个并查集才用了抽象的思维，
       把每一个小块抽象为了4块，并且按照下面的标号进行标记。
                    3 00000 1
                    33 000 11
                    333 0 111
                    33  2  11
                    3  222  1
       注意这4部分是大小是相同的的!
       如果遇到'/' 那就把1区域和2区域合并，0和3合并。
       遇到'\'就把0和1合并，2和3合并
       遇到' ' 就把所有区域合并。
       每一块结束后，再与周围的进行合并
       最后再计算有多少连通块就可以了。

*/
