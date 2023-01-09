package priority_queue

type PriorityQueue[T interface{}] struct {
	array    []T
	lessFunc func(a, b T) bool // if a < b, return true
}

func New[T interface{}](arr []T, lessFunc func(a, b T) bool) *PriorityQueue[T] {

	t := &PriorityQueue[T]{
		array:    arr,
		lessFunc: lessFunc,
	}

	n := t.len()
	for i := n/2 - 1; i >= 0; i-- {
		t.down(i, n)
	}
	return t
}

func (h *PriorityQueue[T]) len() int {
	return len(h.array)
}

func (h *PriorityQueue[T]) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *PriorityQueue[T]) less(i, j int) bool {
	return h.lessFunc(h.array[i], h.array[j])
}

func (h *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(j, i) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

func (h *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(j, i) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *PriorityQueue[T]) Push(x T) {
	h.array = append(h.array, x)
	h.up(h.len() - 1)
}

func (h *PriorityQueue[T]) Pop() T {
	if h.len() == 0 {
		panic("PriorityQueue[T].Pop() require non-empty")
	}

	n := h.len() - 1
	h.swap(0, n)
	h.down(0, n)

	a := h.array[n]
	h.array = h.array[:n]
	return a
}
