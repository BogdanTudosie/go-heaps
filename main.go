package main

import "fmt"

// define a struct to hold the heap (it's actually an array)
type MaxHeap struct {
	array []int
}

// insert element into the array within the Heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swapElements(parent(index), index)
		index = parent((index))
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChildIndex(index int) int {
	return 2*index + 2
}

func rightChildIndex(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swapElements(firstIndex int, secondIndex int) {
	h.array[firstIndex], h.array[secondIndex] = h.array[secondIndex], h.array[firstIndex]
}

func main() {
	heap := &MaxHeap{}
	fmt.Println(heap)
	numberHeap := []int{10, 20, 30, 40}
	for _, value := range numberHeap {
		heap.Insert(value)
	}
	fmt.Println(heap.array)

}
