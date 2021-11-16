package binary_heap

type BinaryHeap struct {
	Data []int
}

func (h *BinaryHeap) Size() int {
	return len(h.Data)
}

func (h *BinaryHeap) heapify(index int) {
	if index > h.Size() {
		return
	}
	top := h.Data[index]
	left := index*2 + 1
	right := index*2 + 2
	if left < h.Size() && h.Data[left] > top {
		h.Data[left], h.Data[index] = h.Data[index], h.Data[left]
	}
	if right < h.Size() && h.Data[right] > top {
		h.Data[right], h.Data[index] = h.Data[index], h.Data[right]
	}
}
