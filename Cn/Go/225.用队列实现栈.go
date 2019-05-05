type MyQueque struct {
	queque []int
}

func QueueConstructor(size int) MyQueque {
	return MyQueque{queque: make([]int, 0, size)}
}

func (q *MyQueque) Enqueque(x int) {
	q.queque = append(q.queque, x)
}

func (q *MyQueque) Dequeque() int {
	var result = q.queque[0]
	q.queque = q.queque[1:len(q.queque)]
	return result
}

func (q *MyQueque) Size() int {
	return len(q.queque)
}

func (q *MyQueque) Front() int {
	return q.queque[0]
}

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */
type MyStack struct {
	a MyQueque
	b MyQueque
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{a: QueueConstructor(100), b: QueueConstructor(100)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	var temp *MyQueque
	if this.b.Size() > 0 {
		temp = &this.b
	} else {
		temp = &this.a
	}

	temp.Enqueque(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	var temp *MyQueque
	var other *MyQueque
	if this.b.Size() > 0 {
		temp = &this.b
		other = &this.a
	} else {
		temp = &this.a
		other = &this.b
	}

	tempLen := temp.Size()
	for i := 0; i < tempLen-1; i++ {
		other.Enqueque(temp.Dequeque())
	}
	result := temp.Dequeque()
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	var temp *MyQueque
	if this.b.Size() > 0 {
		temp = &this.b
	} else {
		temp = &this.a
	}

	return temp.queque[temp.Size()-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.a.Size() == 0 && this.b.Size() == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

