package main

import "fmt"

/**
// test struct初始化
type ListNode struct {
	Val  int
	Next *ListNode
}
func main() {
	a := ListNode{0, nil}
	fmt.Println(a.Next)
}
*/

// test 切片某些特性

func main() {
	slice := make([]int, 3)
	slice[0] = 3
	slice[1] = 2
	slice[2] = 1
	fmt.Printf("%v\n", slice)
	slice = slice[0:1]
	slice[0] = 9
	fmt.Printf("%v\n", slice)
	// fmt.Printf("%v\n", slice[1])
	slice = slice[0:0]
	slice = append(slice, 8)
	slice = append(slice, 7)
	fmt.Printf("%v\n", slice)
	fmt.Printf("%d, %d\n", len(slice), cap(slice))

	slice = slice[1:3]
	fmt.Printf("%v\n", slice)
	slice = slice[1:3]
	fmt.Printf("%v\n", slice)
}