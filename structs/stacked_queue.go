package structs

type StackedQueue struct {
	stack     Stack
	tempStack Stack
}

func NewStackedQueue() *StackedQueue {
	return &StackedQueue{
		stack:     *NewStack(),
		tempStack: *NewStack(),
	}
}

func (q *StackedQueue) Enqueue(x int) {
	for !q.stack.IsEmpty() {
		pop := q.stack.Pop()
		q.tempStack.Push(pop)
	}

	q.stack.Push(x)
	for !q.tempStack.IsEmpty() {
		pop := q.tempStack.Pop()
		q.stack.Push(pop)
	}
}

func (q *StackedQueue) Dequeue() int {
	return q.stack.Pop()
}

func (q *StackedQueue) Peek() int {
	return q.stack.Top()
}

func (q *StackedQueue) IsEmpty() bool {
	return q.stack.IsEmpty()
}
