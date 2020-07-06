package main

import (
	"bytes"
	"fmt"
	"sort"
)

func minNumber(nums []int) string {
	sort.Slice(nums,func(i,t int) bool{
		return splice(nums[i],nums[t])<splice(nums[t],nums[i])
	})
	return stringfy(nums)
}

func splice(A,B int) string{
	return fmt.Sprintf("%d%d",A,B)
}
func stringfy(nums []int) string{
	buffer := bytes.Buffer{}
	for _,num := range nums{
		buffer.WriteString(fmt.Sprintf("%d",num))
	}
	return buffer.String()
}
