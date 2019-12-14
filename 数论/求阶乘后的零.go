package main

// 迭代 求 n! 尾部有多少个零
func trailingZeroes(n int) int {

	ans := 0
	for n != 0 {
		ans += n / 5
		n = n / 5
	}
	return ans
}
/*
	题目链接: https://leetcode-cn.com/problems/factorial-trailing-zeroes/comments/
*/

/*
	总结
	1.  要求n!有多少个零，就是求n!由多少个10相乘。
	    那么由于10由质因子2和5组成，那么我们只需要求n!有多少个因子2和因子5就可以了，再取二者的最小值就可以了。
		又由于n!中，2的因子一定多于5个因子，所以我们只需要算5的因子。 (小伙伴可以验证下)
	2.  如何算n!中5的因子呢？
			我们注意到 [1,5]之间只有1个5，而[1,10]之间有2个5(5贡献1个，10贡献一个)
			[1,30]里面有7个5。 (5,10,15,20,30分别贡献1个，而25贡献2个)
			于是我们可以当做 (5,10,15,20,25,30先贡献1个，之后25再贡献1个)
			于是就可以得到 n!中5个因子 = n/5 + n/25 + n / 125 + .... + n / (5^k)	(k>=1,k∈Z)
*/
