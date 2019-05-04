type Stack struct {
	stack []int
}

func StackConstructor(size int) Stack {
	return Stack{stack: make([]int, 0, size)}
}

func (s *Stack) Push(i int) {
	s.stack = append(s.stack, i)
}

func (s *Stack) Pop() int {
	var stackLen = len(s.stack)
	if stackLen == 0 {
		return 0
	}

	var result = s.stack[stackLen-1]
	s.stack = s.stack[:stackLen-1]
	return result
}

func (s *Stack) Peek() int {
	var stackLen = len(s.stack)
	return s.stack[stackLen-1]
}

func (s *Stack) Len() int {
	return len(s.stack)
}

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
type MyQueue struct {
	a Stack
	b Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{a: StackConstructor(500), b: StackConstructor(500)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.a.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.b.Len() > 0 {
		return this.b.Pop()
	} else {
		for this.a.Len() > 0 {
			this.b.Push(this.a.Pop())
		}
		return this.b.Pop()
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.b.Len() > 0 {
		return this.b.Peek()
	} else {
		for this.a.Len() > 0 {
			this.b.Push(this.a.Pop())
		}
		return this.b.Peek()
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.a.Len() <= 0 && this.b.Len() <= 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

