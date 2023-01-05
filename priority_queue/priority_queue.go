package priority_queue

import (
	"container/heap"
)

func T() {
	heap.Init()
}

type Heap[T interface{}] struct {
	arr []T
}

func New[T interface{}]() *Heap[T] {
	return &Heap[T]{arr: []T{}}
}

func (h *Heap[T]) Len() int {
	return len(h.arr)
}

func (h *Heap[T]) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
