package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{0, nil}
	fmt.Println(a.Next)
}
