package queue

type ArrayQueue struct {
	q    []interface{}
	size int
	head int
	tail int
}

/** Initialize your data structure here. */
func Constructor() *ArrayQueue {
	return &ArrayQueue{
		q:    make([]interface{}, 32),
		size: 0,
	}
}

/** Push element x to the back of queue. */
func (this *ArrayQueue) Push(x interface{}) {
	this.enqueue(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *ArrayQueue) Pop() interface{} {
	return this.dequeue()
}

/** Get the front element. */
func (this *ArrayQueue) Peek() interface{} {
	if this.Empty() {
		return nil
	}
	return this.q[this.head]
}

/** Returns whether the queue is empty. */
func (this *ArrayQueue) Empty() bool {
	return this.size == 0
}

func (this *ArrayQueue) enqueue(x interface{}) bool {
	if this.size == len(this.q) {
		return false
	}
	this.size++
	this.q[this.tail] = x
	this.tail++
	if this.tail == len(this.q) {
		this.tail = 0
	}

	return true
}

func (this *ArrayQueue) dequeue() interface{} {
	if this.Empty() {
		return nil
	}
	x := this.q[this.head]
	this.q[this.head] = nil
	this.size--
	this.head++
	if this.head == len(this.q) {
		this.head = 0
	}

	return x
}

/**
 * Your ArrayQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
