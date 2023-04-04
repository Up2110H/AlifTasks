package queue

import "AlifTasks/stack"

type Queue struct {
	a, b stack.Stack
}

func (q *Queue) IsEmpty() (result bool) {
	// TODO: implement this method
	return result
}

func (q *Queue) Enqueue(val int) {
	// TODO: implement this method
}

func (q *Queue) Dequeue() (val int, err error) {
	// TODO: implement this method
	return val, err
}
