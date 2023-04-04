package queue

import (
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	var q Queue

	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", q.IsEmpty(), true)
	}

	q.Push(1)
	if q.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", q.IsEmpty(), false)
	}

	_, err := q.Pop()
	if err != nil {
		t.Errorf("Pop() error = %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", q.IsEmpty(), true)
	}

}

func TestQueue_Pop(t *testing.T) {
	var q Queue
	_, err := q.Pop()
	if err == nil {
		t.Errorf("Pop() error = %v", err)
	}

	for i := 0; i < 10; i++ {
		q.Push(i)
	}

	for i := 0; i < 10; i++ {
		val, err := q.Pop()
		if err != nil {
			t.Errorf("Pop() error = %v", err)
		}
		if val != i {
			t.Errorf("Pop() = %v, want %v", val, i)
		}
	}

	_, err = q.Pop()
	if err == nil {
		t.Errorf("Pop() error = %v", err)
	}

}
