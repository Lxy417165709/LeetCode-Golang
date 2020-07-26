package Geek

func wiggleSort(nums []int)  {
	for i:=1;i<len(nums);i++{
		if i%2==1{
			if nums[i]>nums[i-1]{
				nums[i],nums[i-1] = nums[i-1],nums[i]
			}
		}else{
			if nums[i]<nums[i-1]{
				nums[i],nums[i-1] = nums[i-1],nums[i]
			}
		}
	}
}

/*
	题目链接: https://leetcode-cn.com/problems/peaks-and-valleys-lcci/
	总结
		1. 这题真的没想到...看了题解才知道怎么做
		2. 第一次的时候想到了先排序，再利用另外一个数组进行穿插，但是显然不太合语义
		3. 题解还有另外2种方法，一种是先排序，另外一种是使用暴力(查找最大最小值进行穿插)
		4. 应该可以用位运算简化上面的代码
*/
