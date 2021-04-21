package main

import (
	"container/heap"
	"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

var _ heap.Interface = (*IntHeap)(nil)

// List Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNode Heap(min-heap)
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int               { return len(h) }
func (h ListNodeHeap) Less(i int, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListNodeHeap) Push(x interface{})    { *h = append(*h, x.(*ListNode)) }
func (h *ListNodeHeap) Pop() interface{} {
	item := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return item
}

var _ heap.Interface = (*ListNodeHeap)(nil)
var _ sort.Interface = (*ListNodeHeap)(nil)
