package struct_util

import (
	"fmt"
	"testing"
)

func TestMyHeap(t *testing.T) {
	heap := NewMyHeap(100, func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})
	heap.Push(3)
	heap.Push(2)
	heap.Push(5)
	heap.Push(1)
	heap.Push(2)
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
	fmt.Println(heap.data)
	heap.Push(22)
	heap.Push(50)
	heap.Push(111)
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
	fmt.Println(heap.data)
	fmt.Println(heap.Pop())
}
