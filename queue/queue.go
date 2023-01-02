package queue

import "container/list"

type Queue[T interface{}] struct {
	list list.List
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *Queue[T]) Push(v T) {
	q.list.PushFront(v)
}

func (q *Queue[T]) Front() T {
	e := q.list.Back()
	if e == nil {
		return
	}
	return
}
