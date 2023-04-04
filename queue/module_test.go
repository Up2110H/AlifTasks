package queue

import (
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	var q Queue

	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", q.IsEmpty(), true)
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", q.IsEmpty(), false)
	}

	_, err := q.Dequeue()
	if err != nil {
		t.Errorf("Dequeue() error = %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", q.IsEmpty(), true)
	}

}

func TestQueue_Dequeue(t *testing.T) {
	var q Queue
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Dequeue() error = %v", err)
	}

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 10; i++ {
		val, err := q.Dequeue()
		if err != nil {
			t.Errorf("Dequeue() error = %v", err)
		}
		if val != i {
			t.Errorf("Dequeue() = %v, want %v", val, i)
		}
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Errorf("Dequeue() error = %v", err)
	}

}
