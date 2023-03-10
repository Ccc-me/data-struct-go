package queue

import (
	"github.com/Ccc-me/data-struct-go/list"
)

type Queue[T interface{}] struct {
	list list.List[T]
}

func New[T interface{}]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Push(v T) {
	q.list.PushBack(v)
}

// 长度为0的情况下，panic
func (q *Queue[T]) Front() T {
	if q.list.Len() == 0 {
		panic("Queue[T].Front() require non-empty")
	}
	return q.list.Front().Value
}

func (q *Queue[T]) Back() T {
	if q.list.Len() == 0 {
		panic("Queue[T].Front() require non-empty")
	}
	return q.list.Back().Value
}

func (q *Queue[T]) Pop() {
	if q.list.Len() == 0 {
		return
	}
	q.list.Remove(q.list.Front())
}
